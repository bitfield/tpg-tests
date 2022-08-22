package script_test

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestScript(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
