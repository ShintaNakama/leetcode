package main

import "testing"

func Test_rob2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				nums: []int{2, 3, 2},
			},
			want: 3,
		},
		{
			name: "no2",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "no3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "no4",
			args: args{
				nums: []int{1, 2, 1, 1},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("%q. rob() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
