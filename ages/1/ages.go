package ages

func Total(data map[string]int) int {
	total := 0
	for _, age := range data {
		total += age
	}
	return total
}
