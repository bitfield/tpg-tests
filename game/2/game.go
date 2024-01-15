package game

import "strings"

func ListItems(items []string) string {
	result := "You can see here"
	result += strings.Join(items, ", ")
	result += "."
	return result
}
