package service

import (
	"net/http"
)

// serviceHandler
func (h *HTTP) svcHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, Magoo"
	RespondJSON(w, msg)
}

// svcListHandler a simple hello request
func (h *HTTP) svcListHandler(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, h)
}
