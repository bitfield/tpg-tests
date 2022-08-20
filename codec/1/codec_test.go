package codec_test

import (
	"math/rand"
	"testing"

	"codec"

	"github.com/google/go-cmp/cmp"
)

func TestEncodeFollowedByDecodeGivesStartingValue(t *testing.T) {
	t.Parallel()
	input := rand.Intn(10)
	encoded := codec.Encode(input)
	t.Logf("encoded value: %#v", encoded)
	// after the round trip, we should get what we started with
	want := input
	got := codec.Decode(encoded)
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
