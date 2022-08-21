package serve3_test

import (
	"net"
	"testing"
	"time"

	serve "serve3"
)

func TestListenAsyncPatiently(t *testing.T) {
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
