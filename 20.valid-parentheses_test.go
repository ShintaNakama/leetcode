package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				s: "(])",
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				s: "(([]){})",
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				s: "[({(())}[()])]",
			},
			want: true,
		},
		{
			name: "9",
			args: args{
				s: "(){}}{",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := isValid(tt.args.s); got != tt.want {
			t.Errorf("%q. isValid() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
