package main

/**
 * <p>Given the <code>head</code> of a sorted linked list, <em>delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list</em>. Return <em>the linked list <strong>sorted</strong> as well</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/linkedlist1.jpg" style="width: 500px; height: 142px;" />
<pre>
<strong>Input:</strong> head = [1,2,3,3,4,4,5]
<strong>Output:</strong> [1,2,5]
</pre>

<p><strong class="example">Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/linkedlist2.jpg" style="width: 500px; height: 205px;" />
<pre>
<strong>Input:</strong> head = [1,1,1,2,3]
<strong>Output:</strong> [2,3]
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
 * [1,2,3,3,4,4,5]
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head2 := deleteDuplicates2(head.Next.Next)
		if head2 != nil && head.Val == head2.Val {
			return head2.Next
		} else {
			head = head2
		}
	} else {
		head.Next = deleteDuplicates2(head.Next)
	}

	return head
}

/*
リンクリストにおいて「重複したノードを全て削除」するという問題を解決するには、特に「前」のノードの情報を保持しておくことが必要です。これは、リンクリストが一方向であるため、後から前のノードに戻ることができないからです。したがって、「前」のノードを追跡することで、重複したノードを削除した後にリンクリストを再接続することが可能となります。

前のノードを追跡するには、通常は「ダミーノード」を使ったり、2つのポインタ（1つは現在のノードを指し、もう1つは前のノードを指す）を使用します。そして、リンクリストを通過する間、これらのポインタを適切に更新することで問題を解決します。

*/

func deleteDuplicates3(head *ListNode) *ListNode {
	// Dummy node will be a "prev" node for head
	// ダミーノードを作成し重複を全て削除するのでheadの前のノードを保持するために使用する
	dummy := &ListNode{0, head}
	// prevは前のノードでユニークなノードとなる
	prev := dummy

	for head != nil {
		// If this is a beginning of duplicates sublist
		// skip to the end of duplicates sublist
		// ノードが重複するかチェック
		if head.Next != nil && head.Val == head.Next.Val {
			// 重複は全て削除するので、head.Nextがnilになるか重複しないノードまでheadを進める
			// Skip to the end of duplicates sublist
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// Skip duplicates
			// ここで前のユニークなノードと重複を全て除いたノードをリンクさせる
			// このタイミングのheadは重複したノードの最後のノードが入っているためhead.Nextを使用する
			prev.Next = head.Next
		} else {
			// Otherwise, move predecessor
			// 重複していない場合はprevを進める
			prev = prev.Next
		}
		// Move forward
		// headを進めて次のノードの重複をチェックする
		head = head.Next
	}

	// dummy.Nextがユニークなノードの先頭となる
	return dummy.Next
}
