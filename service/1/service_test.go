package service_test

import (
	"testing"

	"github.com/bitfield/service"
)

func TestRunningIsTrueWhenServiceIsRunning(t *testing.T) {
	t.Parallel()
	service.Start()
	if !service.Running() {
		t.Error(false)
	}
}
