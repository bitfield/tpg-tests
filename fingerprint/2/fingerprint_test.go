package fingerprint2_test

import (
	"testing"

	fingerprint "fingerprint2"

	"github.com/google/go-cmp/cmp"
)

func TestFingerprint(t *testing.T) {
	t.Parallel()
	data := []byte("These pretzels are making me thirsty.")
	orig := fingerprint.Hash(data)
	same := fingerprint.Hash(data)
	different := fingerprint.Hash([]byte("Hello, Newman"))
	if !cmp.Equal(orig, same) {
		t.Error("same data produced different hash")
	}
	if cmp.Equal(orig, different) {
		t.Error("different data produced same hash")
	}
}
