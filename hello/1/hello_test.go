package hello_test

import (
	"testing"

	"github.com/bitfield/hello"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"hello": hello.Main,
	})
}

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
