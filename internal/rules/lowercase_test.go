package rules

import "testing"

func TestLowercase_Check(t *testing.T) {
	t.Parallel()

	rule := Lowercase{}

	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "valid lowercase",
			msg:  "user created",
			want: "",
		},
		{
			name: "empty message",
			msg:  "",
			want: "",
		},
		{
			name: "starts with uppercase",
			msg:  "User created",
			want: "log message should start with a lowercase letter",
		},
		{
			name: "starts with digit",
			msg:  "1 user created",
			want: "",
		},
		{
			name: "starts with bracket",
			msg:  "[AUTH] user created",
			want: "",
		},
		{
			name: "uppercase in the middle",
			msg:  "add new user ID",
			want: "",
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
