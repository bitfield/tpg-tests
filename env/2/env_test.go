package env2_test

import (
	"net"
	"os"
	"testing"

	env "env2"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"listen": env.Main,
	}))
}

func TestScriptWithExtraEnvVars(t *testing.T) {
	t.Parallel()
	addr := randomLocalAddr(t)
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
		Setup: func(env *testscript.Env) error {
			env.Setenv("SERVER_ADDR", addr)
			return nil
		},
	})
}

func randomLocalAddr(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().String()
}
