package load_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/bitfield/load"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLoadTester_SetsCorrectConfig(t *testing.T) {
	t.Parallel()
	want := &load.LoadTester{
		URL:        "https://example.com",
		Output:     os.Stdout,
		ErrOutput:  os.Stderr,
		HTTPClient: http.DefaultClient,
		Fanout:     10,
		UserAgent:  "loadtest",
	}
	got := load.NewLoadTester("https://example.com", os.Stdout, os.Stderr, http.DefaultClient, 10, "loadtest")
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(os.File{})) {
		t.Error(cmp.Diff(want, got, cmpopts.IgnoreUnexported(os.File{})))
	}
}
