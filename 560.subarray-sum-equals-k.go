package main

/**
 * <p>Given an array of integers <code>nums</code> and an integer <code>k</code>, return <em>the total number of subarrays whose sum equals to</em> <code>k</code>.</p>

<p>A subarray is a contiguous <strong>non-empty</strong> sequence of elements within an array.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> nums = [1,1,1], k = 2
<strong>Output:</strong> 2
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> nums = [1,2,3], k = 3
<strong>Output:</strong> 2
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>-1000 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>-10<sup>7</sup> &lt;= k &lt;= 10<sup>7</sup></code></li>
</ul>

**/
/**
 * [1,1,1]
2
**/

func subarraySum(nums []int, k int) int {
	// mapで部分配列の合計値をkeyにし、その合計値の出現回数を値にする
	m := map[int]int{0: 1}

	var (
		sum   int
		count int
	)

	for _, num := range nums {
		// 要素を加算していく
		sum += num
		// sum-kがmapに存在していれば部分配列の和==kが存在するとみなせる
		// sum kを持つ部分配列があり、その部分配列は累積和が計算された場所の端から始まり、現在いる番号で終わる。
		count += m[sum-k]
		// 合計値をkeyにmapに追加することで計算結果を保存
		m[sum]++

		// fmt.Println("sum", sum)
		// fmt.Println("count", count)
		// fmt.Println("map", m)
	}

	return count
}
