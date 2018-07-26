package service

import (
	"github.com/gorilla/mux"
	log "github.com/rustyeddy/logrus"
)

/*
  HTTP implementation of the Service interface
*/

// HTTP is a single service providing one or more routes
type HTTP struct {
	ID
	running     bool // service currently running?
	*mux.Router      // Pointer to the router for this service
}

// NewHTTP will create a new http service to start and listen to
func NewHTTP(host string, port int) *HTTP {
	s := &HTTP{}
	s.ID = ID{host, port, "http", "http service"}
	s.Router = mux.NewRouter()
	s.running = false
	return s
}

// Ident return the service ID
func (s *HTTP) Ident() *ID {
	return &s.ID
}

// Name is the name of the service
func (s *HTTP) Name() string {
	return s.ID.Name()
}

//Host is the host name of the server this service is running on
func (s *HTTP) Host() string {
	return s.ID.Host()
}

// Port is the port the service is running on
func (s *HTTP) Port() int {
	return s.ID.Port()
}

// Start will start the service
func (s *HTTP) Start() error {
	if !s.running {
		log.Debug("start running service ", s.Name())
	}
	return nil
}

// Stop this service
func (s *HTTP) Stop() error {
	if s.running {
		log.Debug("start the service ", s.Name)
	}
	return nil
}

// Status of this service
func (s *HTTP) Status() string {
	return "what is up?"
}
