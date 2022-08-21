package greet3_test

import (
	"bytes"
	"strings"
	"testing"

	greet "greet3"

	"github.com/google/go-cmp/cmp"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	input := strings.NewReader("fakename\n")
	greet.Greet(input, buf)
	want := "Your name? Hello, fakename!\n"
	got := buf.String()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
