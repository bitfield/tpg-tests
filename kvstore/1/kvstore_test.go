package kvstore_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/bitfield/kvstore"
)

func TestSmokeKVStore(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := range 1000 {
			kvstore.Set("foo", strconv.Itoa(i))
		}
		wg.Done()
	}()
	for i := range 1000 {
		_ = kvstore.Get("foo")
		runtime.Gosched()
	}
	wg.Wait()
}
