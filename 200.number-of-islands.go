package main

import (
	"fmt"
)

/**
 * <p>Given an <code>m x n</code> 2D binary grid <code>grid</code> which represents a map of <code>&#39;1&#39;</code>s (land) and <code>&#39;0&#39;</code>s (water), return <em>the number of islands</em>.</p>

<p>An <strong>island</strong> is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> grid = [
  [&quot;1&quot;,&quot;1&quot;,&quot;1&quot;,&quot;1&quot;,&quot;0&quot;],
  [&quot;1&quot;,&quot;1&quot;,&quot;0&quot;,&quot;1&quot;,&quot;0&quot;],
  [&quot;1&quot;,&quot;1&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;],
  [&quot;0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;]
]
<strong>Output:</strong> 1
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> grid = [
  [&quot;1&quot;,&quot;1&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;],
  [&quot;1&quot;,&quot;1&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;],
  [&quot;0&quot;,&quot;0&quot;,&quot;1&quot;,&quot;0&quot;,&quot;0&quot;],
  [&quot;0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;1&quot;,&quot;1&quot;]
]
<strong>Output:</strong> 3
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 300</code></li>
	<li><code>grid[i][j]</code> is <code>&#39;0&#39;</code> or <code>&#39;1&#39;</code>.</li>
</ul>

**/
/**
 * [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
 */

func numIslands(grid [][]byte) int {
	fmt.Println(grid)
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
				// bfs(grid, i, j)
				fmt.Println(grid)
			}
		}
	}
	fmt.Println(grid)

	return count
}

// DFS(深さ優先探索)なのでgrid[i:len(grid)]を先に確認していく
func dfs(grid [][]byte, r, c int) {
	// gridの範囲外もしくは'1'でない場合はreturn
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != '1' {
		return
	}

	// 再帰でgrid[i:len(grid)]を先に確認していく
	grid[r][c] = '2'
	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}

func bfs(grid [][]byte, r, c int) {
	q := [][2]int{}

	q = append(q, [2]int{r, c})
	grid[r][c] = '2'

	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, offset := range offsets {
			x := curr[0] + offset[0]
			y := curr[1] + offset[1]

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == '1' {
				q = append(q, [2]int{x, y})
				grid[x][y] = 2
			}
		}
	}
}
