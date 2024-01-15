package add_test

import (
	"testing"

	"github.com/bitfield/add"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	add.Add(2, 2)
}
