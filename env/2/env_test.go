package env_test

import (
	"net"
	"testing"

	"github.com/bitfield/env"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"listen": env.Main,
	})
}

func TestWithExtraEnvVars(t *testing.T) {
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
