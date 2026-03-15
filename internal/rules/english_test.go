package rules

import "testing"

func TestEnglish_Check(t *testing.T) {
	t.Parallel()

	rule := English{}

	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "only english",
			msg:  "user created",
			want: "",
		},
		{
			name: "english with digits",
			msg:  "user 123 created",
			want: "",
		},
		{
			name: "cyrillic",
			msg:  "запуск сервера",
			want: "log message should contain only english letters",
		},
		{
			name: "mixed english and cyrillic",
			msg:  "user создан",
			want: "log message should contain only english letters",
		},
		{
			name: "non english latin letter",
			msg:  "user créé",
			want: "log message should contain only english letters",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := rule.Check(test.msg)
			if got != test.want {
				t.Errorf("Check(%q) = %q, want %q", test.msg, got, test.want)
			}
		})
	}
}
