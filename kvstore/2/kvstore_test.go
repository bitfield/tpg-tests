package kvstore2_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	kvstore "kvstore2"
)

func TestSmokeKVStore(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	wg.Add(1)
	s := kvstore.NewStore()
	go func() {
		for i := 0; i < 1000; i++ {
			s.Set("foo", strconv.Itoa(i))
		}
		wg.Done()
	}()
	for i := 0; i < 1000; i++ {
		_ = s.Get("foo")
		runtime.Gosched()
	}
	wg.Wait()
}
