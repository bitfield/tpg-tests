package serve2_test

import (
	"net"
	"testing"
	"time"

	serve "serve2"
)

func TestListenAsyncWithSleep(t *testing.T) {
	t.Parallel()
	addr := randomLocalAddr(t)
	serve.ListenAsync(addr)
	time.Sleep(10 * time.Millisecond)
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
