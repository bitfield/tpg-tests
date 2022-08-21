package add_test

import (
	"testing"

	"add"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	add.Add(2, 2)
}
