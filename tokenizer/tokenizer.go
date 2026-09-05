package tokenizer

import (
	"strings"
)

type token []rune

type tokenizer struct {
	currentIdx int
	target     []rune
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

func (t *tokenizer) Next() token {
	t.ignoreWhitespace()

	builder := strings.Builder{}
	for t.canContinue() {
		if isWhiteSpace(t.GetRune()) {
			break
		}
		builder.WriteRune(t.GetRune())

		t.NextChar()
	}

	return token(builder.String())
}

func (t *tokenizer) canContinue() bool {
	return t.currentIdx < len(t.target)
}

func (t *tokenizer) ignoreWhitespace() {
	for t.canContinue() && isWhiteSpace(t.GetRune()) {
		t.NextChar()
	}
}
