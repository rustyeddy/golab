package store

import (
	"net/http"
	// github.com/rustyeddy/magoo/server
	srv "github.com/rustyeddy/magoo/service"
)

// Handle root requests
func (s *Store) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		srv.RespondErrorJSON(w, 400, "method "+r.Method+" not supported ")
	} else {
		srv.RespondHTML(w, "Hello, World!")
	}
}
