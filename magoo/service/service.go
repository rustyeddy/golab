package service

import "github.com/gorilla/mux"

// Service is a well known provider of a particular type
type Service interface {
	Ident() ID      // Return our service ID
	Start() error   // Start a service pass back a channel
	Stop() error    // Stope the selected service
	Status() string // Expect JSON, but could be any string
	Routes(r *mux.Router)
}

// Router accepts routes and processes messages
type Router interface {
	Routes(r *mux.Router)
}
