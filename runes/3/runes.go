package runes3

import "unicode/utf8"

func FirstRune(s string) rune {
	if s == "" {
		return utf8.RuneError
	}
	return rune(s[0])
}
