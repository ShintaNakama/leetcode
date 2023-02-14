package main

import "math"

/**
 * <p>Given the <code>root</code> of a binary tree, <em>determine if it is a valid binary search tree (BST)</em>.</p>

<p>A <strong>valid BST</strong> is defined as follows:</p>

<ul>
	<li>The left <span data-keyword="subtree">subtree</span> of a node contains only nodes with keys <strong>less than</strong> the node&#39;s key.</li>
	<li>The right subtree of a node contains only nodes with keys <strong>greater than</strong> the node&#39;s key.</li>
	<li>Both the left and right subtrees must also be binary search trees.</li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg" style="width: 302px; height: 182px;" />
<pre>
<strong>Input:</strong> root = [2,1,3]
<strong>Output:</strong> true
</pre>

<p><strong class="example">Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg" style="width: 422px; height: 292px;" />
<pre>
<strong>Input:</strong> root = [5,1,4,null,null,3,6]
<strong>Output:</strong> false
<strong>Explanation:</strong> The root node&#39;s value is 5 but its right child&#39;s value is 4.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the tree is in the range <code>[1, 10<sup>4</sup>]</code>.</li>
	<li><code>-2<sup>31</sup> &lt;= Node.val &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

**/
/**
 * [2,1,3]
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
引数が有効なBST(二分探索木)か判定する
有効なBSTは以下のように定義される。

左側のサブツリーには、そのノードのキーより小さいキーを持つノードのみが含まれる。
ノードの右サブツリーには、そのノードのキーより大きいキーを持つノードのみが含まれる。
左サブツリーと右サブツリーの両方が二分探索木である必要がある。

上記を考慮してDFSを使って左側と右側をそれぞれ確認していく
この時、ノードのキーより左側は小さく右側は大きいことを確認するために、
再帰処理する関数ではノードとその親のキーも引数にするところがポイント
**/
func isValidBST(root *TreeNode) bool {
	return isValidBSTNode(root, math.MinInt, math.MaxInt)
}

func isValidBSTNode(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}

	if node.Val > lower && node.Val < upper {
		// 再帰処理の2回目以降は再帰の前stepでlowerとupperが設定される
		// ex: test_code no.4
		// step1 isValidBSTNode(4, minInt, 5) && isValidBSTNode(6, 5, maxInt)
		// 左側は階層が1つしかないので終了(true)
		// 右側node:6,lower:5,upper:maxInt
		// step2 isValidBSTNode(3, 5, 6) && isValidBSTNode(7, 5, maxInt)
		// 右側node:3,lower:5,upper:6
		// node(3)がlower(5)をよりも小さいためfalse
		return isValidBSTNode(node.Left, lower, node.Val) && isValidBSTNode(node.Right, node.Val, upper)
	}

	return false
}
