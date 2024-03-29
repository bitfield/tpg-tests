package user_test

import (
	"fmt"

	"github.com/bitfield/user"
)

func ExampleUser() {
	u := user.User{Name: "Gopher"}
	fmt.Println(u)
	// Output:
	// {Gopher}
}

func ExampleUser_NameString() {
	u := user.User{Name: "Gopher"}
	fmt.Println(u.NameString())
	// Output:
	// Gopher
}
