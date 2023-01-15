package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
		{
			name: "10",
			args: args{
				s: "[",
			},
			want: false,
		},
		{
			name: "11",
			args: args{
				s: "((",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.args.s)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. isValid() = %v, want %v\n diff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
