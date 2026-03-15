package rules

import "testing"

func TestNoSensitiveData_Check(t *testing.T) {
	t.Parallel()

	rule := NewNoSensitiveData()

	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "valid message",
			msg:  "user authenticated successfully",
			want: "",
		},
		{
			name: "valid token message",
			msg:  "token validated",
			want: "",
		},
		{
			name: "valid api message",
			msg:  "api request completed",
			want: "",
		},
		{
			name: "password with colon",
			msg:  "user password: ",
			want: "log message should not contain sensitive data",
		},
		{
			name: "api_key with equals",
			msg:  "api_key=",
			want: "log message should not contain sensitive data",
		},
		{
			name: "secret with equals and space",
			msg:  "secret =",
			want: "log message should not contain sensitive data",
		},
		{
			name: "token with colonand space",
			msg:  "token : ",
			want: "log message should not contain sensitive data",
		},
		{
			name: "case insensitive",
			msg:  "Password=",
			want: "log message should not contain sensitive data",
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
