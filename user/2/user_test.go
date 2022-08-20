package user_test

import (
	"errors"
	"testing"

	"user"
)

func TestFindUser_GivesErrUserNotFoundForBogusUser(t *testing.T) {
	_, err := user.FindUser("bogus user")
	if !errors.Is(err, user.ErrUserNotFound) {
		t.Errorf("wrong error: %v", err)
	}
}
