package main

import "fmt"

/**
 * <p>There is a robot on an <code>m x n</code> grid. The robot is initially located at the <strong>top-left corner</strong> (i.e., <code>grid[0][0]</code>). The robot tries to move to the <strong>bottom-right corner</strong> (i.e., <code>grid[m - 1][n - 1]</code>). The robot can only move either down or right at any point in time.</p>

<p>Given the two integers <code>m</code> and <code>n</code>, return <em>the number of possible unique paths that the robot can take to reach the bottom-right corner</em>.</p>

<p>The test cases are generated so that the answer will be less than or equal to <code>2 * 10<sup>9</sup></code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<img src="https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png" style="width: 400px; height: 183px;" />
<pre>
<strong>Input:</strong> m = 3, n = 7
<strong>Output:</strong> 28
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> m = 3, n = 2
<strong>Output:</strong> 3
<strong>Explanation:</strong> From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -&gt; Down -&gt; Down
2. Down -&gt; Down -&gt; Right
3. Down -&gt; Right -&gt; Down
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= m, n &lt;= 100</code></li>
</ul>

**/
/**
 * 3
7
**/

// https://leetcode.com/problems/unique-paths/solutions/474165/python-js-java-go-c-by-o-mn-dp-with-explanation/?orderBy=most_votes&languageTags=golang

// dp[i][j]を、ロボットがgrid[i][j]にいるときに右下隅に到達するためのユニークな経路の数とします。
// dp[0][0] = 1 とし、dp[i][0]とdp[0][j]をそれぞれ1と初期化します。これは、ロボットが一番左の列にいる場合、または一番上の行にいる場合に、それぞれ1通りの方法で右下に到達できることを示します。
// dp[i][j] = dp[i-1][j] + dp[i][j-1]です。これは、ロボットが現在のセルに到達する前に、左のセルから来るか、上のセルから来るかを考慮した合計です。
// 最終的な答えは、dp[m-1][n-1]に保存されています。

func uniquePaths(m int, n int) int {
	path := make([][]int, m)

	for i := 0; i < m; i++ {
		row := make([]int, n)
		path[i] = row
	}

	//fmt.Println(path)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				path[i][j] = 1
				continue
			}

			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	fmt.Println(path)

	return path[m-1][n-1]
}
