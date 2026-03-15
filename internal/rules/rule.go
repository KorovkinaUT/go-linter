package rules

type Rule interface {
	Check(msg string) string
}

func DefaultRules() []Rule {
	return []Rule{
		Lowercase{},
		English{},
		NoSpecialSymbols{},
		NewNoSensitiveData(),
	}
}
