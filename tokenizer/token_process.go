package tokenizer

import (
	"errors"
	"strings"
)

var ErrInvalidToken = errors.New("invalid token")

// Possible combination with variable token
func (t *tokenizer) ignoreWhitespaceCharacters() {
	// look at the prefix to get an idea what type of white space is this
	// H41100: sequence of one or more WHITESPACE character
	// H41110: comment with -- beginning
	// 441120: comment with /* */
	if !t.canContinue() {
		return
	}

	// H41100: sequence of one or more WHITESPACE character
	if isWhiteSpace(t.GetRune()) {
		for t.canContinue() && isWhiteSpace(t.GetRune()) {
			t.NextChar()
		}
		return
	}

	// H41110: comment with -- beginning
	if t.GetRune() == '-' && (t.canPeek() && t.peekRune() == '-') {
		t.NextChar()
		t.NextChar()

		for t.canContinue() {
			r := t.GetRune()
			if r == '\n' {
				break
			}

			t.NextChar()
		}

		return
	}

	// 441120: comment with /* */
	if t.GetRune() == '/' && (t.canPeek() && t.peekRune() == '*') {
		t.NextChar()
		t.NextChar()

		for t.canContinue() {
			r := t.GetRune()
			if r == '*' && (t.canPeek() && t.peekRune() == '/') {
				break
			}
			t.NextChar()
		}

		return
	}
}

func (t *tokenizer) processTokenWithAlphabetStart() (Token, error) {
	result := Token{}

	i := 0
	builder := strings.Builder{}
	for t.canContinue() && !isWhiteSpace(t.GetRune()) {
		if i == 0 {
			if !isAlphabetic(t.GetRune()) {
				// first character need to be an alphabet
				return Token{}, ErrInvalidToken
			}

			builder.WriteRune(t.GetRune())
			i++
			t.NextChar()
			continue
		} else {
			if !isAlphanumeric(t.GetRune()) {
				break
			}

			builder.WriteRune(t.GetRune())
			i++
			t.NextChar()
			continue
		}
	}

	result.Type = IDENT
	result.Literal = TokenLiteral(builder.String())

	tokenType, isKeyword := findKeyword(result.Literal.String())
	if isKeyword {
		result.Type = tokenType
	}
	return result, nil
}

// TBD: maybe better to use map, instead of pure logic branches
func (t *tokenizer) processTokenWithSpecialStart() (Token, error) {
	builder := strings.Builder{}
	switch t.GetRune() {
	case '*':
		builder.WriteRune(t.GetRune())
		t.NextChar()
		return Token{STAR, TokenLiteral(builder.String())}, nil
	case ';':
		builder.WriteRune(t.GetRune())
		t.NextChar()
		return Token{SEMI, TokenLiteral(builder.String())}, nil
	case '(':
		builder.WriteRune(t.GetRune())
		t.NextChar()
		return Token{LP, TokenLiteral(builder.String())}, nil
	case ')':
		builder.WriteRune(t.GetRune())
		t.NextChar()
		return Token{RP, TokenLiteral(builder.String())}, nil
	case ',':
		builder.WriteRune(t.GetRune())
		t.NextChar()
		return Token{COMMA, TokenLiteral(builder.String())}, nil
		// TODO: other operator token
	// string literal process
	case '\'':
		// TODO: maybe refactor to its own function later
		builder.WriteRune(t.GetRune())
		t.NextChar()
		for t.GetRune() != '\'' {
			builder.WriteRune(t.GetRune())
			t.NextChar()
		}
		builder.WriteRune(t.GetRune())
		t.NextChar()
		return Token{STRLIT, TokenLiteral(builder.String())}, nil
	}
	return Token{}, ErrInvalidToken
}
