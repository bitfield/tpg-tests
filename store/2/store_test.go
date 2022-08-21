package store2_test

import (
	store "store2"
	"testing"
)

func TestOpenGivesErrUnopenableForBogusFile(t *testing.T) {
	t.Parallel()
	_, err := store.Open("bogus")
	if err != store.ErrUnopenable {
		t.Errorf("wrong error: %v", err)
	}
}
