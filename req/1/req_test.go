package req_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitfield/req"
)

func TestRequestReturnsErrRateLimitWhenRatelimited(t *testing.T) {
	t.Parallel()
	ts := newRateLimitingServer()
	defer ts.Close()
	err := req.Request(ts.URL)
	if !errors.Is(err, req.ErrRateLimit) {
		t.Errorf("wrong error: %v", err)
	}
}

func newRateLimitingServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTooManyRequests)
		}))
}
