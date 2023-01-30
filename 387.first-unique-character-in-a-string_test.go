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
		{
			name: "no.2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "no.3",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstUniqChar(tt.args.s)
			if got != tt.want {
				t.Errorf("%q. firstUniqChar() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
