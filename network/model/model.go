package model

import "net"

type Iface struct {
	Name string
	net.IP
	net.IPNet
	*Host
	*Link
}
