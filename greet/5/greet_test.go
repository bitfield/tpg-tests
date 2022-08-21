package greet5_test

import (
	"bytes"
	"strings"
	"testing"

	greet "greet5"

	"github.com/google/go-cmp/cmp"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	input := strings.NewReader("Mary Jo\n")
	greet.Greet(input, buf)
	want := "Your name? Hello, Mary Jo!\n"
	got := buf.String()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
