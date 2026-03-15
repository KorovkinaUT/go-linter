package rules

import "unicode"

type NoSpecialSymbols struct{}

func (NoSpecialSymbols) Check(msg string) string {
	for _, r := range msg {

		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r == '%' {
			continue
		}

		return "log message should not contain special symbols or emoji"
	}

	return ""
}
