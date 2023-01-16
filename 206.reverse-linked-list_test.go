package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. reverseList() = %v, want %v\ndiff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
