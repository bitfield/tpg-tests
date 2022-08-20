package reader_test

import (
	"reader"
	"strings"
	"testing"
)

func TestReadAll_ReturnsAnyReadError(t *testing.T) {
	t.Parallel()
	input := strings.NewReader("any old data")
	_, err := reader.ReadAll(input)
	if err == nil {
		t.Error("want error for broken reader, got nil")
	}
}
