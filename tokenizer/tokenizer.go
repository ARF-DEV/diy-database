package tokenizer

import (
	"strings"
)

type tokenizer struct {
	currentIdx int
	target     []rune
	lastToken  TokenLiteral
}

func NewTokenizer(target []rune) tokenizer {
	return tokenizer{
		target: []rune(strings.TrimSpace(string(target))),
	}
}

func (t *tokenizer) NextChar() {
	t.currentIdx++
}

func (t *tokenizer) GetRune() rune {
	return t.target[t.currentIdx]
}

func (t *tokenizer) Scan() (scannedToken TokenLiteral) {
	defer func() {
		t.lastToken = scannedToken
	}()

	t.ignoreWhitespace()

	builder := strings.Builder{}
	for t.canContinue() {
		if isWhiteSpace(t.GetRune()) ||
			// a convenient way to break a token that is joint together with semicolon,
			// as of the current implementation i haven't implement a proper way to detect token adn token types
			// so this will be refactored in the future when i implement a proper type detection code
			(len(builder.String()) > 0 && isSemicolon(t.GetRune())) {
			break
		}

		builder.WriteRune(t.GetRune())

		t.NextChar()
	}

	scannedToken = TokenLiteral(builder.String())
	// todo: do a proper token type detection
	if string(scannedToken) == "" {
		scannedToken = TokenLiteral(";")
	}
	return
}

func (t *tokenizer) Next() bool {
	t.ignoreWhitespace()
	// todo: do a proper token type detection
	return string(t.lastToken) != string([]rune(";"))
}

func (t *tokenizer) canContinue() bool {
	return t.currentIdx < len(t.target)
}

func (t *tokenizer) ignoreWhitespace() {
	for t.canContinue() && isWhiteSpace(t.GetRune()) {
		t.NextChar()
	}
}
