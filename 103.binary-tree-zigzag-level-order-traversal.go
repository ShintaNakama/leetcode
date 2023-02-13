package main

/**
 * <p>Given the <code>root</code> of a binary tree, return <em>the zigzag level order traversal of its nodes&#39; values</em>. (i.e., from left to right, then right to left for the next level and alternate between).</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" style="width: 277px; height: 302px;" />
<pre>
<strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> [[3],[20,9],[15,7]]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> root = [1]
<strong>Output:</strong> [[1]]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> root = []
<strong>Output:</strong> []
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the tree is in the range <code>[0, 2000]</code>.</li>
	<li><code>-100 &lt;= Node.val &lt;= 100</code></li>
</ul>

**/
/**
 * [3,9,20,null,null,15,7]
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}

	q := []*TreeNode{
		root,
	}

	zigzagFlag := false

	// level(Treeの階層)ごとに処理する必要があるのでBFS
	// levelごとに結果に返す値を逆転させる必要があるため、
	// 現在のレベルが逆転するlevelかどうかを判別するフラグを定義し
	// levelが進むごとにフラグを変更している
	for len(q) > 0 {
		length := len(q)
		level := make([]int, length)
		for i := 0; i < length; i++ {
			n := q[0]
			q = q[1:]
			if zigzagFlag {
				level[len(level)-1-i] = n.Val
			} else {
				level[i] = n.Val
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		if zigzagFlag {
			zigzagFlag = false
		} else {
			zigzagFlag = true
		}

		res = append(res, level)
	}

	return res
}
