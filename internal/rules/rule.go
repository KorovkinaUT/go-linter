package rules

type Rule interface {
	Description() string
	Check(msg string) string
}

func DefaultRules() []Rule {
	return []Rule{
		Lowercase{},
		English{},
	}
}
