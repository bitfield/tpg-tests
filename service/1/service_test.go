package service_test

import (
	"service"
	"testing"
)

func TestRunningIsTrueWhenServiceIsRunning(t *testing.T) {
	t.Parallel()
	service.Start()
	if !service.Running() {
		t.Error(false)
	}
}
