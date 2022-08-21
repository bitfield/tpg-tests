package game8

import "strings"

func ListItems(items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return "You can see " + items[0] + " here."
	case 2:
		return "You can see here " + items[0] + " and " +
			items[1] + "."
	default:
		return "You can see here " +
			strings.Join(items[:len(items)-1], ", ") +
			", and " + items[len(items)-1] + "."
	}
}
