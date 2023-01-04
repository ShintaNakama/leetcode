package main

import "testing"

func Test_hasCycle(t *testing.T) {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n1,
	}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: n2,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := hasCycle(tt.args.head); got != tt.want {
			t.Errorf("%q. hasCycle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
