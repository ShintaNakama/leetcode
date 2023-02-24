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
func uniquePaths(m int, n int) int {

	// Create 2D array with size m x n
	// pathはsize m*n で2Dを再現している
	// rowをpathをフラットにしたもの
	path := make([][]int, m)
	rows := make([]int, m*n)
	for y := 0; y < m; y++ {
		path[y] = rows[y*n : (y+1)*n]
	}
	fmt.Println(path)
	fmt.Println(rows)

	// ループの最後でpathの合計がpath[m-1][n-1]に代入されるようになっている
	// まだ理解できていない
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {

			//fmt.Println(path)
			if (y == 0) || (x == 0) {
				path[y][x] = 1
			} else {
				// just follow the DP recurrence formula
				path[y][x] = path[y-1][x] + path[y][x-1]
			}
			fmt.Println(path)

		}
	}

	fmt.Println(path)

	return path[m-1][n-1]
}
