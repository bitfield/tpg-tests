package load_test

import (
	"bytes"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/bitfield/load"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLoadTester(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	want := &load.LoadTester{
		URL:        "https://example.com",
		Output:     buf,
		ErrOutput:  buf,
		HTTPClient: &http.Client{Timeout: time.Second},
		Fanout:     20,
		UserAgent:  "loadtest",
	}
	got := load.NewLoadTester("https://example.com",
		load.WithOutput(buf),
		load.WithErrOutput(buf),
		load.WithHTTPClient(&http.Client{Timeout: time.Second}),
		load.WithConcurrentRequests(20),
		load.WithUserAgent("loadtest"),
	)
	ignore := cmpopts.IgnoreUnexported(os.File{}, bytes.Buffer{})
	if !cmp.Equal(want, got, ignore) {
		t.Error(cmp.Diff(want, got, ignore))
	}
}
