package rules

import "unicode"

type English struct{}

func (English) Check(msg string) string {
	for _, r := range msg {

		if unicode.IsLetter(r) &&
			!(r >= 'a' && r <= 'z') &&
			!(r >= 'A' && r <= 'Z') {

			return "log message should contain only english letters"
		}
	}

	return ""
}
