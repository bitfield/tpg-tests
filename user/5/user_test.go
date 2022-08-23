package user5_test

import (
	"testing"

	user "user5"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	user.Create("Alice")
	if !user.Exists("Alice") {
		t.Error("Alice not created")
	}
}
