package store_test

import (
	"store"
	"testing"
)

func TestOpenGivesErrUnopenableForBogusFile(t *testing.T) {
	_, err := store.Open("bogus")
	if err != store.ErrUnopenable {
		t.Errorf("wrong error: %v", err)
	}
}
