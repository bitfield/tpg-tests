package user7_test

import (
	"testing"

	user "user7"
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
	createUserOrFail(t, "Alice")
	createUserOrFail(t, "Bob")
	user.Delete("Alice")
	if user.Exists("Alice") {
		t.Error("Alice still exists after delete")
	}
	if !user.Exists("Bob") {
		t.Error("Bob was unexpectedly deleted")
	}
}

func createUserOrFail(t *testing.T, name string) {
	t.Helper()
	user.Create(name)
	if !user.Exists(name) {
		t.Errorf("%s not created", name)
	}
}
