package game5

import "strings"

func ListItems(items []string) string {
	result := "You can see here "
	if len(items) < 3 {
		return result + items[0] + " and " + items[1] + "."
	}
	result += strings.Join(items[:len(items)-1], ", ")
	result += ", and "
	result += items[len(items)-1]
	result += "."
	return result
}
