package tokenizer

import (
	"strings"
)

type tokenizer struct {
	currentIdx  int
	target      []rune
	lastToken   TokenLiteral
	lastTokenV2 Token
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

func (t *tokenizer) ScanV2() (scannedToken Token) {
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

		if isSpecial(t.GetRune()) {
			// possibility:
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
	t.lastTokenV2 = scannedToken
	return
}

func (t *tokenizer) canPeek() bool {
	return t.currentIdx+1 < len(t.target)
}

func (t *tokenizer) peekRune() rune {
	return t.target[t.currentIdx+1]
}

func (t *tokenizer) Next() bool {
	t.ignoreWhitespace()
	// todo: do a proper token type detection
	return string(t.lastToken) != string([]rune(";"))
}

func (t *tokenizer) NextV2() bool {
	return t.lastTokenV2.Type != SEMI
}

func (t *tokenizer) canContinue() bool {
	return t.currentIdx < len(t.target)
}

func (t *tokenizer) ignoreWhitespace() {
	for t.canContinue() && isWhiteSpace(t.GetRune()) {
		t.NextChar()
	}
}
