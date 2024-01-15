package batch_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bitfield/batch"
)

func TestRunBatchJob(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Millisecond)
	defer cancel()
	go func() {
		batch.RunBatchJob(ctx)
		cancel()
	}()
	<-ctx.Done()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		t.Fatal("timed out")
	}
}
