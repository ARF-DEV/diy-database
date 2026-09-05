package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	t.Run("tokenizer shall process from left to right", func(t *testing.T) {
		tknizer := NewTokenizer([]rune("hello world"))
		processedString := ""
		for tknizer.canContinue() {
			processedString += string(tknizer.GetRune())
			tknizer.NextChar()
		}

		assert.Equal(t, "hello world", processedString)
	})

	t.Run("whitespace token only acts as seperator between tokens", func(t *testing.T) {
		tokenizer := NewTokenizer([]rune("Hello world"))

		expected := []token{[]rune("Hello"), []rune("world")}
		tokens := []token{}

		for {
			if !tokenizer.canContinue() {
				break
			}

			token := tokenizer.Next()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens)
	})

	t.Run("whitespace token only acts as seperator between tokens -> multiple whitespace", func(t *testing.T) {
		tokenizer := NewTokenizer([]rune("Hello      world    "))

		expected := []token{[]rune("Hello"), []rune("world")}
		tokens := []token{}

		for {
			if !tokenizer.canContinue() {
				break
			}

			token := tokenizer.Next()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens)
	})
}
