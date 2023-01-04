package main

import "testing"

func Test_addRecurcive(t *testing.T) {
	type args struct {
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
				n: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addRecurcive(tt.args.n); got != tt.want {
				t.Errorf("addRecurcive() = %v, want %v", got, tt.want)
			}
		})
	}
}
