package greet7_test

import (
	"os"
	"testing"

	greet "greet7"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"greet": greet.Main,
	}))
}

func TestGreet(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
