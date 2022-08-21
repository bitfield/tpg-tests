package valid4_test

import (
	"testing"
	valid "valid4"
)

func TestValidIsTrueForValidInput(t *testing.T) {
	t.Parallel()
	if !valid.Valid("valid input") {
		t.Error(false)
	}
}

func TestValidIsFalseForInvalidInput(t *testing.T) {
	t.Parallel()
	if valid.Valid("invalid input") {
		t.Error(true)
	}
}
