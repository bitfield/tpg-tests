package fingerprint_test

import (
	"crypto/md5"
	"testing"

	"github.com/bitfield/fingerprint"
)

func TestHashReturnsMD5HashOfGivenData(t *testing.T) {
	t.Parallel()
	data := []byte("These pretzels are making me thirsty.")
	want := md5.Sum(data)
	got := fingerprint.Hash(data)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
