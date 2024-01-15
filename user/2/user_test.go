package user_test

import (
	"errors"
	"testing"

	"github.com/bitfield/user"
)

func TestFindUser_GivesErrUserNotFoundForBogusUser(t *testing.T) {
	t.Parallel()
	_, err := user.FindUser("bogus user")
	if !errors.Is(err, user.ErrUserNotFound) {
		t.Errorf("wrong error: %v", err)
	}
}
