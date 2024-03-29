package main

import "fmt"

/**
 * <p>Given the <code>head</code> of a sorted linked list, <em>delete all duplicates such that each element appears only once</em>. Return <em>the linked list <strong>sorted</strong> as well</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/list1.jpg" style="width: 302px; height: 242px;" />
<pre>
<strong>Input:</strong> head = [1,1,2]
<strong>Output:</strong> [1,2]
</pre>

<p><strong class="example">Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/list2.jpg" style="width: 542px; height: 222px;" />
<pre>
<strong>Input:</strong> head = [1,1,2,3,3]
<strong>Output:</strong> [1,2,3]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the list is in the range <code>[0, 300]</code>.</li>
	<li><code>-100 &lt;= Node.val &lt;= 100</code></li>
	<li>The list is guaranteed to be <strong>sorted</strong> in ascending order.</li>
</ul>

**/
/**
 * [1,1,2]
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//// 自力解答
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	if head.Val == head.Next.Val {
//		head = deleteDuplicates(head.Next)
//	} else {
//		head.Next = deleteDuplicates(head.Next)
//	}
//
//	return head
//}

// 以下のような解答もある
// Valが重複した場合は重複削除のためにNextをNext.Nextで上書き(重複削除)する
// Val重複していない場合はNextとNext.Nextが重複していないか確認するためにcurrentをNextで入れ替える
// currentはheadのcopyなのでheadを返す
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		fmt.Println(current)
		fmt.Println(head)
		if current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}

	return head
}
