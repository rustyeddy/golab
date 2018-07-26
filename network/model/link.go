/* A Link is a network connection that may be share be two (or more)
   hosts, depending on network type.  The link may represent a
   physical (ethernet) netowrk or a logical (GRE Tunnel) network, your
   choice.

   Links facilitate the connection between hosts.

   Host1.IFace1 -> link1 <- 2ecafI.2tsoH (Host2.Iface2 backwards)
*/
package model

import "net"

type Link struct {
	name   string
	ifaces Hashmap
	net.IPNet
}
