package magoo

import (
	"encoding/json"
	"fmt"
	. "github.com/rustyeddy/golib"
	"log"
	//"math/rand"
)

/*

Entry is a File type, hence it also has the respective entries:
- Name
- Path
- ... and so on
*/

// Entry is a response from a submit
type Entry struct {
	Name  string `json:"Name"`  // Unique name for this entry
	Form  string `json:"Form"`  // Form name
	Group string `json:"Group"` // FormGroup

	// Possibly for every entry
	Values    map[string]string `json:"Values"`    // Field values all string
	Timestamp string            `json:"Timestamp"` // timestamp yyyymmdd-hhmm
}

// NewEntry creates a new but empty entry
func NewEntry() *Entry {
	return &Entry{
		Name: GetNewEntryName(),
	}
}

// NewEntryVars creates a new entry from a set of variables
func EntryFromArgs(args *Args) (e *Entry) {
	e = &Entry{
		Form:   args.GetDefault("form", "anon"),
		Group:  args.GetDefault("group", ""),
		Values: args.ToMap(),
	}
	e.Name = GetNewEntryName()
	return e
}

// GetNewEntryName create a new and unique entry name
func GetNewEntryName() string {
	ts := GetTimeStamp()
	return fmt.Sprintf("e-%s", ts[11:16])
}

// JSON
func (e *Entry) JSON() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		log.Printf("failed to marshal JSON")
		return nil
	}
	return b
}
