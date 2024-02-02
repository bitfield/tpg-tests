package serve_test

import (
	"net"
	"testing"
	"time"

	"github.com/bitfield/serve"
)

func TestListenAsync_ListensEventually(t *testing.T) {
	t.Parallel()
	addr := randomLocalAddr(t)
	serve.ListenAsync(addr)
	_, err := net.Dial("tcp", addr)
	for err != nil {
		t.Log("retrying")
		time.Sleep(time.Millisecond)
		_, err = net.Dial("tcp", addr)
	}
}

func randomLocalAddr(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().String()
}
