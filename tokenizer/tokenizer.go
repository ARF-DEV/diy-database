package tokenizer

// Tokenizer follows Requirements For The SQLite Tokenizer

type charClass byte

const (
	WHITESPACE   charClass = 0
	ALPHABETIC   charClass = 1
	NUMERIC      charClass = 2
	ALPHANUMERIC charClass = 3
	HEXADECIMAL  charClass = 4
	SPECIAL      charClass = 5
)

// One of these five characters: u0009, u000a, u000c, u000d, or u0020
func isWhiteSpace(character rune) bool {
	if character == 0x009 ||
		character == 0x000a ||
		character == 0x000c ||
		character == 0x000d ||
		character == 0x0020 {
		return true
	}
	return false
}

// Any of the characters in the range u0041 through u005a (letters "A" through "Z")
// or in the range u0061 through u007a (letters "a" through "z")
// or the character u005f ("_") or any other character larger than u007f.
func isAlphabetic(character rune) bool {
	if character >= 'A' && character <= 'Z' ||
		character >= 'a' && character <= 'z' ||
		character == '_' ||
		character > 0x007f {
		return true
	}
	return false
}

// Any of the characters in the range u0030 through u0039 (digits "0" through "9")
func isNumeric(character rune) bool {
	if character >= '0' && character <= '9' {
		return true
	}
	return false
}

// Any character which is either ALPHABETIC or NUMERIC
func isAlphanumeric(character rune) bool {
	return isAlphabetic(character) || isNumeric(character)
}

// type tokenizer struct{}
//
// func (t *tokenizer)  {
// }
