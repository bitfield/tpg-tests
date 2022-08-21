package even2

var cache = map[int]bool{}

func IsEven(n int) (even bool) {
	even = n%2 == 0
	if _, ok := cache[n]; !ok {
		cache[n] = n%2 == 0
		even = n%2 == 0
	}
	return even
}
