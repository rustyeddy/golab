package model

import "net"

type Interface struct {
	*Link     // link we are attacted to
	*Host     // host we belong to
	net.IPNet // IP address and subnet we are a member of
}
