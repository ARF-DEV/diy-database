package tokenizer

import (
	"strings"
)

type tokenizer struct {
	currentIdx int
	target     []rune
	lastToken  Token
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

func (t *tokenizer) Scan() (scannedToken Token) {
	var err error
	t.ignoreWhitespaceCharacters()

	for t.canContinue() {
		// TODO: other tokens detection
		if isAlphabetic(t.GetRune()) {
			// possibility:
			// ident
			// keyword
			scannedToken, err = t.processTokenWithAlphabetStart()
			if err != nil {
				panic(err)
			}
			break
		}
		if isNumeric(t.GetRune()) {
			// TODO: maybe refactor to its own function later
			builder := strings.Builder{}
			for t.canContinue() && isNumeric(t.GetRune()) {
				builder.WriteRune(t.GetRune())
				t.NextChar()
			}
			scannedToken = Token{Type: INTLIT, Literal: TokenLiteral(builder.String())}
			break
		}

		if isSpecial(t.GetRune()) {
			// possibility:
			// literals
			// operator
			// variable
			scannedToken, err = t.processTokenWithSpecialStart()
			if err != nil {
				panic(err)
			}
			break
		}

		// skipp, there is still some syntax that isn't implemented yet
		t.NextChar()
	}

	if !t.canContinue() && scannedToken.Type == "" {
		scannedToken = Token{
			Type:    SEMI,
			Literal: TokenLiteral(";"),
		}
	}
	t.lastToken = scannedToken
	return
}

func (t *tokenizer) canPeek() bool {
	return t.currentIdx+1 < len(t.target)
}

func (t *tokenizer) peekRune() rune {
	return t.target[t.currentIdx+1]
}

func (t *tokenizer) Next() bool {
	return t.lastToken.Type != SEMI
}

func (t *tokenizer) canContinue() bool {
	return t.currentIdx < len(t.target)
}
