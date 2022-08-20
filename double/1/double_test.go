package double_test

import (
	"double"
	"testing"
)

func TestDouble2Returns4(t *testing.T) {
	t.Parallel()
	want := 4
	got := double.Double(2)
	if want != got {
		t.Errorf("Double(2): want %d, got %d", want, got)
	}
}
