package model

// Node is an addressable element on the Network
type Node struct {
	Hostname string
	Domain   string
	Ifmgmt   *Iface
	Ifaces   Hashmap // Additional interfaces
}

type Host struct {
	Node
}

type Server struct {
	Node
	Services []*Service
}
