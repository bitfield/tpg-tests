package service_test

import (
	"service"
	"testing"
)

func TestRunningIsTrueWhenServiceIsRunning(t *testing.T) {
	service.Start()
	if !service.Running() {
		t.Error(false)
	}
}
