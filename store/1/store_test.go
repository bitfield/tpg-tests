package store_test

import (
	"errors"
	"store"
	"testing"
)

func TestOpenGivesNonNilErrorForBogusFile(t *testing.T) {
	_, err := store.Open("bogus file")
	if err == nil {
		t.Error("want error opening bogus store file")
	}
}

func TestOpenGivesSpecificErrorForBogusFile(t *testing.T) {
	want := errors.New("open bogus: no such file or directory")
	_, got := store.Open("bogus")
	if got != want {
		t.Errorf("wrong error: %v", got)
	}
}

func TestOpenGivesSpecificErrorStringForBogusFile(t *testing.T) {
	want := errors.New("open bogus: no such file or directory")
	_, got := store.Open("bogus")
	if got.Error() != want.Error() {
		t.Errorf("wrong error: %v", got)
	}
}

func TestOneErrorValueEqualsAnother(t *testing.T) {
	want := errors.New("Go home, Go, you're drunk")
	got := errors.New("Go home, Go, you're drunk")
	if got != want {
		t.Errorf("wrong error: %v", got)
	}
}
