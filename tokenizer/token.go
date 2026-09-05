// Package tokenizer
package tokenizer

type TokenType string

const (
	Whitespace TokenType = "Whitespace"
	Ident      TokenType = "Identifier"
	Lit        TokenType = "Literal"
	Var        TokenType = "Variable"
	Op         TokenType = "Operator"
	Keyword    TokenType = "Keyword"
)
