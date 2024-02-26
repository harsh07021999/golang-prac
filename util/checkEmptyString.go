package util

import "unicode"

func CheckEmptyOrBlankString(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
