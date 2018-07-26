package service

import (
	"github.com/gorilla/mux"
)

// Routes will install routes on the given router
func (s *HTTP) Routes(r *mux.Router) {
	r.HandleFunc("/", s.svcHandler)
}
