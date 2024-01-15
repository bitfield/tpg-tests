package serve_test

import (
	"net"
	"testing"

	"github.com/bitfield/serve"
)

func TestListenAsyncEagerly(t *testing.T) {
	t.Parallel()
	addr := randomLocalAddr(t)
	serve.ListenAsync(addr)
	_, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
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
