package user_test

import (
	"testing"

	"github.com/bitfield/user"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	if user.Exists("Alice") {
		t.Fatal("Alice unexpectedly exists")
	}
	user.Create("Alice")
	if !user.Exists("Alice") {
		t.Error("Alice not created")
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	user.Create("Alice")
	if !user.Exists("Alice") {
		t.Error("Alice not created")
	}
	user.Delete("Alice")
	if user.Exists("Alice") {
		t.Error("Alice still exists after delete")
	}
}
