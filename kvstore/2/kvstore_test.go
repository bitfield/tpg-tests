package kvstore_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/bitfield/kvstore"
)

func TestKVStore_HasNoDataRace(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	wg.Add(1)
	s := kvstore.NewStore()
	go func() {
		for i := range 1000 {
			s.Set("foo", strconv.Itoa(i))
		}
		wg.Done()
	}()
	for range 1000 {
		_ = s.Get("foo")
		runtime.Gosched()
	}
	wg.Wait()
}
