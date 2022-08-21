package user3_test

import (
	"errors"
	"testing"

	user "user3"
)

func TestFindUser_GivesErrUserNotFoundForBogusUser(t *testing.T) {
	t.Parallel()
	_, err := user.FindUser("bogus user")
	if !errors.As(err, &user.ErrUserNotFound{}) {
		t.Errorf("wrong error: %v", err)
	}
}
