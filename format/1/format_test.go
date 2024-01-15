package format_test

import (
	"testing"

	"github.com/bitfield/format"
)

const (
	invalidInput        = "invalid input"
	validInput          = "valid input"
	validInputFormatted = "valid input formatted"
)

func TestFormatData_ErrorsOnInvalidInput(t *testing.T) {
	t.Parallel()
	_, err := format.Data(invalidInput)
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestFormatData_IsCorrectForValidInput(t *testing.T) {
	t.Parallel()
	want := validInputFormatted
	got, err := format.Data(validInput)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
