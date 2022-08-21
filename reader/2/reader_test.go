package reader2_test

import (
	"io"
	reader "reader2"
	"testing"
)

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestReadAll_ReturnsAnyReadError(t *testing.T) {
	t.Parallel()
	input := errReader{} // always returns error
	_, err := reader.ReadAll(input)
	if err == nil {
		t.Error("want error for broken reader, got nil")
	}
}
