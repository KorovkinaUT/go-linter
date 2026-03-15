package rules

import "strings"

type NoSensitiveData struct {
	patterns []string
}

func NewNoSensitiveData() *NoSensitiveData {
	words := []string{
		"password",
		"pwd",
		"token",
		"api_key",
		"api key",
		"apikey",
		"secret",
		"private_key",
		"private key",
		"privatekey",
	}

	patterns := make([]string, 0, len(words)*3)

	for _, word := range words {
		patterns = append(patterns, word+":")
		patterns = append(patterns, word+" :")
		patterns = append(patterns, word+"=")
		patterns = append(patterns, word+" =")
	}

	return &NoSensitiveData{patterns: patterns}
}

func (r *NoSensitiveData) Check(msg string) string {
	lower := strings.ToLower(msg)

	for _, p := range r.patterns {
		if strings.Contains(lower, p) {
			return "log message should not contain sensitive data"
		}
	}

	return ""
}
