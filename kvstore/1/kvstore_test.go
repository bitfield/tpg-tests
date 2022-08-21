package kvstore_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"kvstore"
)

func TestSmokeKVStore(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			kvstore.Set("foo", strconv.Itoa(i))
		}
		wg.Done()
	}()
	for i := 0; i < 1000; i++ {
		_ = kvstore.Get("foo")
		runtime.Gosched()
	}
	wg.Wait()
}
