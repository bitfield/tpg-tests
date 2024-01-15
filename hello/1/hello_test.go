package hello_test

import (
	"os"
	"testing"

	"github.com/bitfield/hello"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"hello": hello.Main,
	}))
}

func TestHello(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
