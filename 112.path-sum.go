package main

/**
 * <p>Given the <code>root</code> of a binary tree and an integer <code>targetSum</code>, return <code>true</code> if the tree has a <strong>root-to-leaf</strong> path such that adding up all the values along the path equals <code>targetSum</code>.</p>

<p>A <strong>leaf</strong> is a node with no children.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg" style="width: 500px; height: 356px;" />
<pre>
<strong>Input:</strong> root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
<strong>Output:</strong> true
<strong>Explanation:</strong> The root-to-leaf path with the target sum is shown.
</pre>

<p><strong class="example">Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
<pre>
<strong>Input:</strong> root = [1,2,3], targetSum = 5
<strong>Output:</strong> false
<strong>Explanation:</strong> There two root-to-leaf paths in the tree:
(1 --&gt; 2): The sum is 3.
(1 --&gt; 3): The sum is 4.
There is no root-to-leaf path with sum = 5.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> root = [], targetSum = 0
<strong>Output:</strong> false
<strong>Explanation:</strong> Since the tree is empty, there are no root-to-leaf paths.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the tree is in the range <code>[0, 5000]</code>.</li>
	<li><code>-1000 &lt;= Node.val &lt;= 1000</code></li>
	<li><code>-1000 &lt;= targetSum &lt;= 1000</code></li>
</ul>

**/
/**
 * [5,4,8,11,null,13,4,7,2,null,null,null,1]
22
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
DFSアルゴリズム(深さ優先探索)のフレームワークを考える

あるノードがいわゆる葉ノードであるのは、以下の場合のみである。
空ノードでなく、左子と右子の両方を持たない場合。

経路和は、目標値を指定してルートからリーフまで宝探しをするようなものだと想像する

目標値を前回の目標値-現在のノード値として更新し、DFSで次のレベルに降りる。
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	// base case
	if root == nil {
		return false
	}

	// targetSumを現在のrootの値との差分を求めて
	// trueの条件(合計がtargetSumと同じでリーフノードの場合)に合致したらtrue
	sub := targetSum - root.Val
	if sub == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	// ここでDFSで次のレベルに降りる
	// つまりrootのLeftとRightを再帰的に探索していく
	return hasPathSum(root.Left, sub) || hasPathSum(root.Right, sub)
}
