package main

/**
 * <p>You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and <b>it will automatically contact the police if two adjacent houses were broken into on the same night</b>.</p>

<p>Given an integer array <code>nums</code> representing the amount of money of each house, return <em>the maximum amount of money you can rob tonight <b>without alerting the police</b></em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,3,1]
<strong>Output:</strong> 4
<strong>Explanation:</strong> Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,7,9,3,1]
<strong>Output:</strong> 12
<strong>Explanation:</strong> Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 400</code></li>
</ul>

**/
/**
 * [1,2,3,1]
**/

/**
この問題は、「隣接する要素の選択を避けつつ、配列内の要素の最大合計を見つける」という一般的な問題であり、動的計画法（DP）を使用して解決できます。

まず、配列の先頭から順番に各要素を選択し、その要素を含む場合と含まない場合の最大値を保持する二つの変数、curr_includeとcurr_excludeを保持します。curr_includeは現在の要素を含む場合の最大値、curr_excludeは現在の要素を含まない場合の最大値です。次の要素に進むたびに、これらの変数を更新します。

curr_includeは、前の要素を含まない場合の最大値(curr_exclude)に現在の要素の値を加えたものになります。curr_excludeは、前の要素を含む場合と含まない場合の最大値のうち、大きい方を選択します。

これらの変数を配列の末尾まで更新することで、最終的な結果が得られます。最後に、curr_includeとcurr_excludeのうち、大きい方を返します。

このアルゴリズムは、配列の各要素を一度だけ処理するため、時間計算量はO(n)です。また、curr_includeとcurr_excludeの二つの変数だけを使用するため、空間計算量はO(1)
**/

func rob(nums []int) int {
	currInclude := 0
	currExclude := 0
	for _, num := range nums {
		temp := currInclude
		currInclude = currExclude + num
		currExclude = max(temp, currExclude)
		//fmt.Println(currInclude)
		//fmt.Println(currExclude)
	}
	return max(currInclude, currExclude)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
