package serve3

import "net"

func ListenAsync(addr string) {
	go net.Listen("tcp", addr)
}
