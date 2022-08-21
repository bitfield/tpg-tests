package even4

import "sync"

var (
	m     = sync.Mutex{}
	cache = map[int]bool{}
)

func IsEven(n int) (even bool) {
	m.Lock()
	defer m.Unlock()
	even, ok := cache[n]
	if !ok {
		even = n%2 == 0
		cache[n] = even
	}
	return even
}
