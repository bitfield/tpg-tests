package user_test

import (
	"testing"

	"user"

	"github.com/google/go-cmp/cmp"
)

func TestGreetingReturnsCorrectGreetingForLanguage(t *testing.T) {
	t.Parallel()
	u := user.New("刘慈欣")
	u.Language = "Chinese"
	want := "你好"
	got := u.Greeting()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
