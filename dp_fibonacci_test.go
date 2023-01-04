package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "n=2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "n=3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "n=4",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "n=5",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib3(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fib(b *testing.B) {
	b.ResetTimer()
	fib(15)
}

func Benchmark_fib2(b *testing.B) {
	b.ResetTimer()
	fib2(100)
}

func Benchmark_fib3(b *testing.B) {
	b.ResetTimer()
	fib3(100)
}
