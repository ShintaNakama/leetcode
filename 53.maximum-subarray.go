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

// https://leetcode.com/problems/maximum-subarray/solutions/2392041/brute-to-kadane/?orderBy=most_votes&languageTags=golang
// この問題では部分配列がsumがmaxになる場合のmaxの値を求める
// maxの値を保持し最大値を更新していけば良い
// sumがマイナスになったら0を代入していい
func maxSubArray(nums []int) int {
	var (
		sum int
		max = nums[0]
	)

	for _, n := range nums {
		sum += n
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}

		//fmt.Println("n", n)
		//fmt.Println("sum", sum)
		//fmt.Println("max", max)

	}
	return max
}
