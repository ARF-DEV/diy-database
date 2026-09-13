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
	IDENT TokenType = "IDENT"
	LIT   TokenType = "LIT"
	VAR   TokenType = "VAR"

	// Keyword
	SELECT TokenType = "KEYWORD_SELECT"
	FROM   TokenType = "KEYWORD_FROM"

	// Operator
	STAR TokenType = "OPERATOR_STAR" // "*"
	SEMI TokenType = "OPERATOR_SEMI" // ";"
)

// supported keywords
var keywordMap map[string]TokenType = map[string]TokenType{
	"SELECT": SELECT,
	"FROM":   FROM,
}
