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
	//
	t.Run("whitespace token only acts as seperator between tokens", func(t *testing.T) {
		tokenizer := NewTokenizer([]rune("Hello world"))

		expected := []Token{
			{
				Type:    IDENT,
				Literal: TokenLiteral("Hello"),
			},
			{
				Type:    IDENT,
				Literal: TokenLiteral("world"),
			},
			{
				Type:    SEMI,
				Literal: TokenLiteral(";"),
			},
		}
		tokens := []Token{}

		for tokenizer.Next() {
			token := tokenizer.Scan()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens, "expected: %v, but got: %v", expected, tokens)
	})

	t.Run("whitespace token only acts as seperator between tokens -> multiple whitespace", func(t *testing.T) {
		tokenizer := NewTokenizer([]rune("Hello      world    "))
		expected := []Token{
			{
				Type:    IDENT,
				Literal: TokenLiteral("Hello"),
			},
			{
				Type:    IDENT,
				Literal: TokenLiteral("world"),
			},
			{
				Type:    SEMI,
				Literal: TokenLiteral(";"),
			},
		}
		tokens := []Token{}

		for tokenizer.Next() {
			token := tokenizer.Scan()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens, "expected: %v, but got: %v", expected, tokens)
	})

	t.Run("tokenizer return semicolon at the end of an input when the input has no semicolon as its last character", func(t *testing.T) {
		input := []rune("SELECT * FROM users")
		tokenizer := NewTokenizer(input)

		expected := []Token{
			{
				Type:    SELECT,
				Literal: TokenLiteral("SELECT"),
			},
			{
				Type:    STAR,
				Literal: TokenLiteral("*"),
			},
			{
				Type:    FROM,
				Literal: TokenLiteral("FROM"),
			},
			{
				Type:    IDENT,
				Literal: TokenLiteral("users"),
			},
			{
				Type:    SEMI,
				Literal: TokenLiteral(";"),
			},
		}
		tokens := []Token{}

		for tokenizer.Next() {
			token := tokenizer.Scan()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens, "expected: %v, but got: %v", expected, tokens)
	})

	t.Run("tokenizer didn't return semicolon at the end of an input when the input has no semicolon as its last character", func(t *testing.T) {
		input := []rune("SELECT * FROM users;")
		tokenizer := NewTokenizer(input)

		expected := []Token{
			{
				Type:    SELECT,
				Literal: TokenLiteral("SELECT"),
			},
			{
				Type:    STAR,
				Literal: TokenLiteral("*"),
			},
			{
				Type:    FROM,
				Literal: TokenLiteral("FROM"),
			},
			{
				Type:    IDENT,
				Literal: TokenLiteral("users"),
			},
			{
				Type:    SEMI,
				Literal: TokenLiteral(";"),
			},
		}
		tokens := []Token{}

		for tokenizer.Next() {
			token := tokenizer.Scan()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens, "expected: %v, but got: %v", expected, tokens)
	})

	t.Run("tokenizer sucessfully process a simple insert query", func(t *testing.T) {
		input := []rune("INSERT INTO description VALUES ('hallo', 1)")
		tokenizer := NewTokenizer(input)

		expected := []Token{
			{
				Type:    INSERT,
				Literal: TokenLiteral("INSERT"),
			},
			{
				Type:    INTO,
				Literal: TokenLiteral("INTO"),
			},
			{
				Type:    IDENT,
				Literal: TokenLiteral("description"),
			},
			{
				Type:    VALUES,
				Literal: TokenLiteral("VALUES"),
			},
			{
				Type:    LP,
				Literal: TokenLiteral("("),
			},
			{
				Type:    STRLIT,
				Literal: TokenLiteral("'hallo'"),
			},
			{
				Type:    COMMA,
				Literal: TokenLiteral(","),
			},
			{
				Type:    INTLIT,
				Literal: TokenLiteral("1"),
			},
			{
				Type:    RP,
				Literal: TokenLiteral(")"),
			},
			{
				Type:    SEMI,
				Literal: TokenLiteral(";"),
			},
		}
		tokens := []Token{}

		for tokenizer.Next() {
			token := tokenizer.Scan()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens, "expected: %v, but got: %v", expected, tokens)
	})

	t.Run("tokenizer sucessfully process a simple insert query all lowercase", func(t *testing.T) {
		input := []rune("insert into description values ('hallo', 1)")
		tokenizer := NewTokenizer(input)

		expected := []Token{
			{
				Type:    INSERT,
				Literal: TokenLiteral("insert"),
			},
			{
				Type:    INTO,
				Literal: TokenLiteral("into"),
			},
			{
				Type:    IDENT,
				Literal: TokenLiteral("description"),
			},
			{
				Type:    VALUES,
				Literal: TokenLiteral("values"),
			},
			{
				Type:    LP,
				Literal: TokenLiteral("("),
			},
			{
				Type:    STRLIT,
				Literal: TokenLiteral("'hallo'"),
			},
			{
				Type:    COMMA,
				Literal: TokenLiteral(","),
			},
			{
				Type:    INTLIT,
				Literal: TokenLiteral("1"),
			},
			{
				Type:    RP,
				Literal: TokenLiteral(")"),
			},
			{
				Type:    SEMI,
				Literal: TokenLiteral(";"),
			},
		}
		tokens := []Token{}

		for tokenizer.Next() {
			token := tokenizer.Scan()
			tokens = append(tokens, token)
		}

		assert.EqualValues(t, expected, tokens, "expected: %v, but got: %v", expected, tokens)
	})
}
