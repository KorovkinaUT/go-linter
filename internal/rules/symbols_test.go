package rules

import "testing"

func TestNoSpecialSymbols_Check(t *testing.T) {
	t.Parallel()

	rule := NoSpecialSymbols{}

	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "valid message",
			msg:  "server started",
			want: "",
		},
		{
			name: "valid with digits",
			msg:  "request 123 completed",
			want: "",
		},
		{
			name: "valid format string",
			msg:  "request %d completed",
			want: "",
		},
		{
			name: "contains exclamation mark",
			msg:  "conection failed!!!",
			want: "log message should not contain special symbols or emoji",
		},
		{
			name: "contains dots",
			msg:  "something went wrong...",
			want: "log message should not contain special symbols or emoji",
		},
		{
			name: "contains colon",
			msg:  "warning: something went wrong",
			want: "log message should not contain special symbols or emoji",
		},
		{
			name: "contains emoji",
			msg:  "server started 🚀",
			want: "log message should not contain special symbols or emoji",
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
