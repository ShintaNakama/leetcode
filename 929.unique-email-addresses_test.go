package main

import (
	"testing"
)

func Test_numUniqueEmails(t *testing.T) {
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no.1",
			args: args{
				emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			},
			want: 2,
		},
		{
			name: "no.2",
			args: args{
				emails: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numUniqueEmails(tt.args.emails); got != tt.want {
				t.Errorf("%q. numUniqueEmails = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
