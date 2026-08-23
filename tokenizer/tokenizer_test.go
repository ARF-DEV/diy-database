package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWhiteSpace(t *testing.T) {
	t.Run("Is a whitespace", func(t *testing.T) {
		assert.Equal(t, true, isWhiteSpace(' '))
		assert.Equal(t, true, isWhiteSpace(0x0020))
		assert.Equal(t, true, isWhiteSpace(0x0009))
		assert.Equal(t, true, isWhiteSpace(0x000a))
		assert.Equal(t, true, isWhiteSpace(0x000c))
		assert.Equal(t, true, isWhiteSpace(0x000d))
	})
	t.Run("Not a whitespace", func(t *testing.T) {
		assert.Equal(t, false, isWhiteSpace('A'))
	})
}

func TestIsAlphabetic(t *testing.T) {
	t.Run("Is an alphabet", func(t *testing.T) {
		assert.Equal(t, true, isAlphabetic('A'))
		assert.Equal(t, true, isAlphabetic('a'))
		assert.Equal(t, true, isAlphabetic('_'))
		assert.Equal(t, true, isAlphabetic('_'))
	})
	t.Run("Not an alphabet", func(t *testing.T) {
		assert.Equal(t, false, isAlphabetic('{'))
	})
}

func TestIsNumeric(t *testing.T) {
	t.Run("Is a numeric", func(t *testing.T) {
		assert.Equal(t, true, isNumeric('2'))
	})

	t.Run("Not a numeric", func(t *testing.T) {
		assert.Equal(t, false, isNumeric('A'))
	})
}
