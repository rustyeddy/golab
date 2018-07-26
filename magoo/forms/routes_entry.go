package server

import (
	//"fmt"
	//"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/rustyeddy/golib"
)

type MagooHeader struct {
	Version string
	Magooid string
	Entries interface{}
}

func NewMagooHeader(o interface{}) *MagooHeader {
	mh := &MagooHeader{
		Version: "magoo-v02",
		Magooid: "doe",
		Entries: o,
	}
	return mh
}

// GetFilemap
func (s *HTTPServer) entryFilemapHandler(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, NewMagooHeader(s.EntryStore().Index()))
}

// entryPathList handler
func (s *HTTPServer) entryPathlistHandler(w http.ResponseWriter, r *http.Request) {
	fm := s.EntryStore().Index()
	var paths []string
	var names []string
	for name, path := range fm {
		paths = append(paths, path)
		names = append(names, name)
	}

	lists := struct {
		Paths []string
		Names []string
	}{paths, names}

	RespondJSON(w, NewMagooHeader(lists))
}

// entryNameList handler
func (s *HTTPServer) entryNamelistHandler(w http.ResponseWriter, r *http.Request) {
	fm := s.EntryStore().Index()
	var names []string
	for name, _ := range fm {
		names = append(names, name)
	}
	RespondJSON(w, NewMagooHeader(names))
}

// entryNameHandler will handle request with templates
func (s *HTTPServer) entryNameHandler(w http.ResponseWriter, r *http.Request) {
	es := s.EntryStore()
	args := mux.Vars(r)
	name, e := args["name"]
	if !e {
		RespondErrorJSON(w, 400, "Must include name for argument")
		return
	}

	// Reading from the file the text should already be JSON, send it as is
	b, err := es.Read(name)
	if err != nil {
		log.Printf("entry.Read() name (%s) read entry error %v ", name, err)
		RespondError(w, 500, "Server error reading object")
		return
	}
	RespondText(w, string(b))
}

// ====================== Functions that do the work ========================
// EntryFromRequest
func entryFromRequest(r *http.Request) *Entry {
	args := ArgsFromRequest(r)
	return EntryFromArgs(args)
}

// ==================== Entry Submit Handler ========================
// entrySubmitHandler recieving and process incoming entry submission from form
// Change this to go directly to EntryStorage (rather than server).
func (s *HTTPServer) entrySubmitHandler(w http.ResponseWriter, r *http.Request) {

	// Extract the arguments from the http request.  The request comes
	// in externally, hence we have no control over what might be in the
	// request forms, make sure we sanitize the input.
	args := ArgsFromRequest(r)
	DieNil(args)

	// Get entry storage
	estore := s.EntryStore()
	DieNil(estore) // GAK

	// Store the entry - the entries name must be unique, hence we need
	// to create new storge for this item.
	entry := EntryFromArgs(args)
	if entry == nil {
		RespondError(w, 500, "failed to create entry")
		return
	}
	err := estore.Create(entry.Name, entry.JSON())
	if err != nil {
		RespondError(w, 500, "failed to create entry")
		return
	}
	RespondJSON(w, entry)
}
