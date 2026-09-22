// Package tokenizer
package tokenizer

import (
	"strings"
)

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
	// LIT   TokenType = "LIT"
	// VAR   TokenType = "VAR"

	// Keywords
	SELECT TokenType = "KEYWORD_SELECT"
	FROM   TokenType = "KEYWORD_FROM"
	INSERT TokenType = "KEYWORD_INSERT"
	INTO   TokenType = "KEYWORD_INTO"
	VALUES TokenType = "KEYWORD_VALUES"
	UPDATE TokenType = "KEYWORD_UPDATE"
	SET    TokenType = "KEYWORD_SET"
	WHERE  TokenType = "KEYWORD_WHERE"

	// Operators
	STAR  TokenType = "OPERATOR_STAR"         // "*"
	SEMI  TokenType = "OPERATOR_SEMI"         // ";"
	LP    TokenType = "OPERATOR_LEFT_PARENT"  // "("
	RP    TokenType = "OPERATOR_RIGHT_PARENT" // ")"
	COMMA TokenType = "OPERATOR_COMMA"        // ","
	EQ    TokenType = "OPERATOR_EQUAL"

	// Literals
	STRLIT TokenType = "LITERAL_STR"
	INTLIT TokenType = "LITERAL_INT"
)

// supported keywords
var keywordMap map[string]TokenType = map[string]TokenType{
	"SELECT": SELECT,
	"FROM":   FROM,
	"INSERT": INSERT,
	"VALUES": VALUES,
	"INTO":   INTO,
	"UPDATE": UPDATE,
	"SET":    SET,
	"WHERE":  WHERE,
}

func findKeyword(str string) (TokenType, bool) {
	tokenType, isKeyword := keywordMap[strings.ToUpper(str)]
	return tokenType, isKeyword
}
