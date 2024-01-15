package codec_test

import (
	"math/rand"
	"testing"

	"github.com/bitfield/codec"
)

func TestEncodeFollowedByDecodeGivesStartingValue(t *testing.T) {
	t.Parallel()
	input := rand.Intn(10) + 1
	encoded := codec.Encode(input)
	t.Logf("encoded value: %#v", encoded)
	want := input
	got := codec.Decode(encoded)
	// after the round trip, we should get what we started with
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
