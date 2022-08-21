package user2_test

import (
	"errors"
	"testing"

	user "user2"
)

func TestFindUser_GivesErrUserNotFoundForBogusUser(t *testing.T) {
	t.Parallel()
	_, err := user.FindUser("bogus user")
	if !errors.Is(err, user.ErrUserNotFound) {
		t.Errorf("wrong error: %v", err)
	}
}
