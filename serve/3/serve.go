package serve

import "net"

func ListenAsync(addr string) {
	go net.Listen("tcp", addr)
}
