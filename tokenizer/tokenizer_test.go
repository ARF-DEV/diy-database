package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizerProcess(t *testing.T) {
	t.Run("tokenizer shall process from left to right", func(t *testing.T) {
		tknizer := NewTokenizer([]rune("hello world"))
		processedString := ""
		for tknizer.canContinue() {
			processedString += string(tknizer.GetRune())
			tknizer.NextChar()
		}

		assert.Equal(t, "hello world", processedString)
	})
}

func TestTokenizer(t *testing.T) {
	cases := []struct {
		name     string
		expected any
		input    []rune
	}{
		{
			name: "whitespace token only acts as seperator between tokens",
			expected: []Token{
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
			},
			input: []rune("Hello world"),
		},
		{
			name: "whitespace token only acts as seperator between tokens -> multiple whitespace",
			expected: []Token{
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
			},
			input: []rune("Hello world"),
		},
		{
			name: "tokenizer return semicolon at the end of an input when the input has no semicolon as its last character",
			expected: []Token{
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
			},
			input: []rune("SELECT * FROM users"),
		},
		{
			name: "tokenizer didn't return semicolon at the end of an input when the input has no semicolon as its last character",
			expected: []Token{
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
			},
			input: []rune("SELECT * FROM users;"),
		},
		{
			name: "tokenizer sucessfully process a simple insert query",
			expected: []Token{
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
					Type:    COMMA,
					Literal: TokenLiteral(","),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("age"),
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
			},
			input: []rune("INSERT INTO description, age VALUES ('hallo', 1);"),
		},
		{
			name: "tokenizer sucessfully process a simple insert query all lowercase",
			expected: []Token{
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
					Type:    COMMA,
					Literal: TokenLiteral(","),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("age"),
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
			},
			input: []rune("insert into description, age values ('hallo', 1)"),
		},
		{
			name: "tokenizer sucessfully process a simple update query",
			expected: []Token{
				{
					Type:    UPDATE,
					Literal: TokenLiteral("UPDATE"),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("users"),
				},
				{
					Type:    SET,
					Literal: TokenLiteral("SET"),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("name"),
				},
				{
					Type:    EQ,
					Literal: TokenLiteral("="),
				},
				{
					Type:    STRLIT,
					Literal: TokenLiteral("'prabs'"),
				},
				{
					Type:    COMMA,
					Literal: TokenLiteral(","),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("age"),
				},
				{
					Type:    EQ,
					Literal: TokenLiteral("="),
				},
				{
					Type:    INTLIT,
					Literal: TokenLiteral("12"),
				},
				{
					Type:    WHERE,
					Literal: TokenLiteral("WHERE"),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("id"),
				},
				{
					Type:    EQ,
					Literal: TokenLiteral("="),
				},
				{
					Type:    INTLIT,
					Literal: TokenLiteral("3"),
				},
				{
					Type:    SEMI,
					Literal: TokenLiteral(";"),
				},
			},
			input: []rune("UPDATE users SET name = 'prabs', age = 12 WHERE id = 3"),
		},
		{
			name: "tokenizer sucessfully process a simple delete query",
			expected: []Token{
				{
					Type:    DELETE,
					Literal: TokenLiteral("DELETE"),
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
					Type:    WHERE,
					Literal: TokenLiteral("where"),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("id"),
				},
				{
					Type:    EQ,
					Literal: TokenLiteral("="),
				},
				{
					Type:    INTLIT,
					Literal: TokenLiteral("10"),
				},
				{
					Type:    SEMI,
					Literal: TokenLiteral(";"),
				},
			},
			input: []rune("DELETE FROM users where id = 10"),
		},
		{
			name: "tokenizer sucessfully process a simple create table query",
			expected: []Token{
				{
					Type:    CREATE,
					Literal: TokenLiteral("CREATE"),
				},
				{
					Type:    TABLE,
					Literal: TokenLiteral("TABLE"),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("users"),
				},
				{
					Type:    LP,
					Literal: TokenLiteral("("),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("id"),
				},
				{
					Type:    INTEGERTYPE,
					Literal: TokenLiteral("INTEGER"),
				},
				{
					Type:    PRIMARY,
					Literal: TokenLiteral("PRIMARY"),
				},
				{
					Type:    KEY,
					Literal: TokenLiteral("KEY"),
				},
				{
					Type:    COMMA,
					Literal: TokenLiteral(","),
				},
				{
					Type:    IDENT,
					Literal: TokenLiteral("name"),
				},
				{
					Type:    STRINGTYPE,
					Literal: TokenLiteral("STRING"),
				},
				{
					Type:    RP,
					Literal: TokenLiteral(")"),
				},
				{
					Type:    SEMI,
					Literal: TokenLiteral(";"),
				},
			},
			input: []rune("CREATE TABLE users (id INTEGER PRIMARY KEY, name STRING)"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tokenizer := NewTokenizer(tc.input)
			tokens := []Token{}

			for tokenizer.Next() {
				token := tokenizer.Scan()
				fmt.Println(token)
				tokens = append(tokens, token)
			}

			assert.Equal(t, tc.expected, tokens)
		})
	}
}
