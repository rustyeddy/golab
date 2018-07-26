package model

import "net"

type Network struct {
	name   string
	domain string
	nodes  Hashmap
	links  Hashmap
	subnet net.IPNet
}
