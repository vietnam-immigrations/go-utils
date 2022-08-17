package text_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/text"
)

func TestRemoveNonASCII(t *testing.T) {
	assert.Equal(t, "L Nam Trng.pdf", text.RemoveNonASCII("Lê Nam Trường.pdf"))
}

func TestSanitizeCVFileName(t *testing.T) {
	sanitized := text.SanitizeCVFileName("Lê. Nam Trường.pdf")
	fmt.Println(sanitized)
	assert.True(t, strings.HasPrefix(sanitized, "L  Nam Trng"))
	assert.True(t, strings.HasSuffix(sanitized, ".pdf"))
}
