package main

/**
 * <p>Given an integer array <code>nums</code>, find the <span data-keyword="subarray-nonempty">subarray</span> with the largest sum, and return <em>its sum</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [-2,1,-3,4,-1,2,1,-5,4]
<strong>Output:</strong> 6
<strong>Explanation:</strong> The subarray [4,-1,2,1] has the largest sum 6.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1]
<strong>Output:</strong> 1
<strong>Explanation:</strong> The subarray [1] has the largest sum 1.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [5,4,-1,7,8]
<strong>Output:</strong> 23
<strong>Explanation:</strong> The subarray [5,4,-1,7,8] has the largest sum 23.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow up:</strong> If you have figured out the <code>O(n)</code> solution, try coding another solution using the <strong>divide and conquer</strong> approach, which is more subtle.</p>

**/
/**
 * [-2,1,-3,4,-1,2,1,-5,4]
**/

// https://leetcode.com/problems/maximum-subarray/solutions/290007/go-4-ms-100-00-easy-code/?orderBy=most_votes&languageTags=golang
// 最大値と合計値を更新していけば良い
func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	//fmt.Println(max, sum)
	for _, v := range nums[1:] {
		// 合計値がマイナスの場合は合計値を更新
		if sum < 0 {
			sum = v
		} else {
			// 合計値がマイナスではない場合は合計値に加算
			sum += v
		}
		// 最大値より合計値の方が大きい場合は最大値を合計値で更新
		// 最大値は常に最大値として扱えるのでnumsを全て参照するO(1)で計算できる
		if max < sum {
			max = sum
		}
		//fmt.Println(max, sum)
	}
	return max
}
