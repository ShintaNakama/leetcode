package main

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "no2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("%q. uniquePaths() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
