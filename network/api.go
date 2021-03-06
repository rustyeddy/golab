/*

The Network package abstracts just enough of a "network" consisting
of hosts and links, where the hosts have one or more interfaces connected
to a single link.

*/
package network

import "net"

// =============== Containers and Maps =================
type Network interface {
	Name() string
	Nodes() *Hashmap
	Links() *Hashmap
	IPSubnet() *net.IPNet
}

type Hashmap interface {
	Get(name string) interface{}
	Set(name string, item interface{}) error
	Exists(name string) bool
	Names() []string       // an array of index names
	Values() []interface{} // values, whatever they may be
}
