package valid_test

import (
	"testing"

	"github.com/bitfield/valid"
)

func TestValid(t *testing.T) {
	t.Parallel()
	want := true
	got := valid.Valid("valid input")
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}
