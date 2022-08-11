package text_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vietnam-immigrations/go-utils/pkg/text"
)

func TestRemoveNonASCII(t *testing.T) {
	assert.Equal(t, "L Nam Trng.pdf", text.RemoveNonASCII("Lê Nam Trường.pdf"))
}
