package service

import (
	"errors"

	"github.com/gorilla/mux"
)

// Server is host (metal, virtual or container) that a service runs on
type Server struct {
	ID
	Router   *mux.Router    // Our router
	Services map[ID]Service // and our service map
	alive    bool           // is it running?
}

// New will create a new server
func New(name, host string) *Server {
	s := &Server{
		ID:       ID{"host", 1199, "magoo", "new description"},
		Services: make(map[ID]Service, 10),
		Router:   mux.NewRouter(),
	}
	return s
}

// Start will start a server on a specific
func (s *Server) Start() error {
	for _, r := range s.Services {
		r.Start()
	}
	return nil
}

// Stop the services
func (s *Server) Stop() error {
	for _, r := range s.Services {
		r.Stop()
	}
	return nil
}

// Add service to sever
func (s *Server) Add(svc Service) error {

	// Get our ID
	id := svc.Ident()

	// Verify that service does not already exist
	if _, ex := s.Services[id]; ex {
		return errors.New("service.Add failed already exists " + id.Name())
	}

	// Register the service on the Services map
	s.Services[id] = svc

	// Now call the svc to install its routes on the router we provide
	svc.Routes(s.Router)

	// That is it, we need to either do something or exit
	return nil
}

// Get service returns the requested service if it exists
func (s *Server) Get(name ID) Service {
	if svc, ex := s.Services[name]; ex {
		return svc
	}
	return nil
}

// Remove the service from manager
func (s *Server) Remove(name ID) error {
	if _, e := s.Services[name]; e {
		delete(s.Services, name)
	}
	return nil
}
