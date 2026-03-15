package rules

import (
	"unicode"
	"unicode/utf8"
)

type Lowercase struct{}

func (Lowercase) Description() string {
	return "first letter in lowercase"
}

func (Lowercase) Check(msg string) string {
	if len(msg) == 0 {
		return ""
	}

	r, _ := utf8.DecodeRuneInString(msg)
	if unicode.IsLetter(r) && !unicode.IsLower(r) {
		return "log message should start with a lowercase letter"
	}

	return ""
}
