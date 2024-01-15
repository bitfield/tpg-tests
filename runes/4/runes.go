package runes

import "unicode/utf8"

func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return utf8.RuneError
}
