package user2

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

var userDB map[string]*User

var ErrUserNotFound = errors.New("user not found")

func FindUser(name string) (*User, error) {
	user, ok := userDB[name]
	if !ok {
		return nil, fmt.Errorf("%q: %w", name, ErrUserNotFound)
	}
	return user, nil
}
