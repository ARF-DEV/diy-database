// Package tokenizer
package tokenizer

type (
	TokenType    string
	TokenLiteral []rune
)

func (t TokenLiteral) String() string {
	return string(t)
}

type Token struct {
	Type    TokenType
	Literal TokenLiteral
}

const (
	Whitespace TokenType = "Whitespace"
	Ident      TokenType = "Identifier"
	Lit        TokenType = "Literal"
	Var        TokenType = "Variable"
	Op         TokenType = "Operator"
	Keyword    TokenType = "Keyword"
)
