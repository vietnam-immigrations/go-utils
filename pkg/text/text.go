package text

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dchest/uniuri"
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

func SanitizeCVFileName(s string) string {
	parts := strings.Split(s, ".")
	if len(parts) < 2 {
		return s
	}
	fileName := RemoveNonASCII(strings.Join(parts[0:len(parts)-1], " "))
	fileType := parts[len(parts)-1]
	return fmt.Sprintf("%s--%s.%s", fileName, uniuri.NewLen(5), fileType)
}
