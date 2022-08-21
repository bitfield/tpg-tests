package even3

import "sync"

var (
	m     = sync.Mutex{}
	cache = map[int]bool{}
)

func IsEven(n int) (even bool) {
	m.Lock()
	defer m.Unlock()
	even = n%2 == 0
	if _, ok := cache[n]; !ok {
		cache[n] = n%2 == 0
		even = n%2 == 0
	}
	return even
}
