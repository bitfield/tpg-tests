package user_test

import (
	"testing"

	"github.com/bitfield/user"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	user.Create("Alice")
	if !user.Exists("Alice") {
		t.Error("Alice not created")
	}
}
