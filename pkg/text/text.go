package text

import (
	"strings"
	"unicode"
)

// RemoveNonASCII removes all non-ascii characters
func RemoveNonASCII(s string) string {
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, s)
}
