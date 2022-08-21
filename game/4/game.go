package game4

import "strings"

func ListItems(items []string) string {
	result := "You can see here "
	result += strings.Join(items[:len(items)-1], ", ")
	result += ", and "
	result += items[len(items)-1]
	result += "."
	return result
}
