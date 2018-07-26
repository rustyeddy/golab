package service

import (
	"strconv"

	log "github.com/rustyeddy/logrus"
)

// ID maintains the tuple that can uniquely identify
// a particular service.
type ID struct {

	// ID
	host string // Host the service runs on
	port int    // Port the service listens on
	name string // Name of the service

	// Meta desciption
	description string
}

// NewID creates a new ID
func NewID(host string, port int, name string, desc string) *ID {
	id := &ID{host, port, name, desc}
	return id
}

// Set ID will set the ID values
func (id *ID) Set(args map[string]string) {
	for k, v := range args {
		switch k {
		case "name":
			id.name = v
		case "host":
			id.host = v
		case "port":
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Error(err.Error())
			}
			id.port = i
		case "desc":
			id.description = v
		default:
			log.Warn("unknown argument ", k)
		}
	}
}

// Name of the service - it is indexable
func (id *ID) Name() string {
	return id.name
}

// Port service is running on
func (id *ID) Port() int {
	return id.port
}

// Host service is running on
func (id *ID) Host() string {
	return id.host
}

// HostPort concats the host and the port for localhost:1199
// format string to be used by the http.ListenAndServe() command
func (id *ID) String() string {
	return id.Host() + ":" + string(id.Port())
}
