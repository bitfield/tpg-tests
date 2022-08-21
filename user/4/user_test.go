package user4_test

import (
	"testing"

	user "user4"

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
