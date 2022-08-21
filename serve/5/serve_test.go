package serve5_test

import (
	"net"
	"testing"
	"time"

	serve "serve5"
)

func TestListenAsyncWithHelper(t *testing.T) {
	t.Parallel()
	addr := randomLocalAddr(t)
	serve.ListenAsync(addr)
	waitForServer(t, addr)
	// test whatever we're really testing
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

func waitForServer(t *testing.T, addr string) {
	t.Helper()
	timeout := time.NewTimer(100 * time.Millisecond)
	_, err := net.Dial("tcp", addr)
	for err != nil {
		select {
		case <-timeout.C:
			t.Fatal("timed out")
		default:
			t.Log("retrying: ", err)
			time.Sleep(time.Millisecond)
			_, err = net.Dial("tcp", addr)
		}
	}
}
