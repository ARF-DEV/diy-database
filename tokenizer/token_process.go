package tokenizer

func (t *tokenizer) processWhiteSpaceToken() (token Token) {
	token.Type = Whitespace
	// look at the prefix to get an idea what type of white space is this
	// H41100: sequence of one or more WHITESPACE character
	// H41110: comment with -- beginning
	// 441120: comment with /* */

	// H41100: sequence of one or more WHITESPACE character
	if isWhiteSpace(t.GetRune()) {
		for t.canContinue() && isWhiteSpace(t.GetRune()) {
			token.Literal = append(token.Literal, t.GetRune())
			t.NextChar()
		}

		return token
	}

	// H41110: comment with -- beginning
	if t.GetRune() == '-' && (t.canPeek() && t.peekRune() == '-') {
		token.Literal = append(token.Literal, t.GetRune())
		t.NextChar()
		token.Literal = append(token.Literal, t.GetRune())
		t.NextChar()

		for t.canContinue() {
			r := t.GetRune()
			if r == '\n' {
				break
			}

			token.Literal = append(token.Literal, r)
			t.NextChar()
		}

		return token
	}

	// 441120: comment with /* */
	if t.GetRune() == '/' && (t.canPeek() && t.peekRune() == '*') {
		token.Literal = append(token.Literal, t.GetRune())
		t.NextChar()
		token.Literal = append(token.Literal, t.GetRune())
		t.NextChar()

		for t.canContinue() {
			r := t.GetRune()
			if r == '*' && (t.canPeek() && t.peekRune() == '/') {
				break
			}
			token.Literal = append(token.Literal, r)
			t.NextChar()
		}

		return token
	}

	return
}
