package greet2_test

import (
	"bytes"
	"testing"

	greet "greet2"

	"github.com/google/go-cmp/cmp"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	greet.Greet(buf)
	want := "Your name? Hello, !\n"
	got := buf.String()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
