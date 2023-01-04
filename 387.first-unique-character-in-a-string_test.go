package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no.1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := firstUniqChar(tt.args.s); got != tt.want {
			t.Errorf("%q. firstUniqChar() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
