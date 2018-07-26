package store

import (
	"github.com/gorilla/mux"
)

// Routes mounts the entry points, I guess you could
// say, getting them ready to be called
func (s *Store) Routes(r *mux.Router) {
	r.HandleFunc("/", s.rootHandler)

	// ============================================================
	// r.HandleFunc("/index/", s.hIndex)
	// r.Handle("/files", s.hFiles)
	// r.Handle("/object/", s.hObjects)
	// r.Handle("/object/{name}", s.hGetObject).Method("GET")
	// r.Handle("/object/{name}", s.hDeleteObject.Method("DELETE"))
	// r.Handle("/object/{post}", s.hPostObject.Method("POST"))
	// =============================================================
}
