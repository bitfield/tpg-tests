package env2

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func Main() int {
	if len(os.Args[1:]) < 1 {
		fmt.Fprintln(os.Stderr, "usage: listen ADDR")
		return 1
	}
	srv := http.Server{
		Addr: os.Args[1],
	}
	done := make(chan struct{})
	go func() {
		<-done
		srv.Shutdown(context.Background())
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the Go web server")
		close(done)
	})
	srv.Handler = mux
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
