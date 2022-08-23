package fingerprint_test

import (
	"crypto/md5"
	"testing"

	"fingerprint"

	"github.com/google/go-cmp/cmp"
)

func TestFingerprint(t *testing.T) {
	t.Parallel()
	data := []byte("These pretzels are making me thirsty.")
	want := md5.Sum(data)
	got := fingerprint.Hash(data)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
