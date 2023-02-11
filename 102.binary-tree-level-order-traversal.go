package main

/**
 * <p>Given the <code>root</code> of a binary tree, return <em>the level order traversal of its nodes&#39; values</em>. (i.e., from left to right, level by level).</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" style="width: 277px; height: 302px;" />
<pre>
<strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> [[3],[9,20],[15,7]]
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
	<li><code>-1000 &lt;= Node.val &lt;= 1000</code></li>
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

// 深さ(level)ごとに値を確認する必要があるのでBFS(幅優先探索)を使う
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}

	q := []*TreeNode{
		root,
	}

	for len(q) > 0 {
		levelSlice := []int{}
		nodes := q
		q = []*TreeNode{}
		for _, n := range nodes {
			levelSlice = append(levelSlice, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		res = append(res, levelSlice)
	}

	return res
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}

	q := []*TreeNode{
		root,
	}

	for len(q) > 0 {
		levelSlice := []int{}
		glen := len(q)
		for i := 0; i < glen; i++ {
			n := q[0]
			levelSlice = append(levelSlice, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			q = q[1:]
		}

		res = append(res, levelSlice)
	}

	return res
}
