package main

/**
 * <p>Given an integer array <code>nums</code>, return <em>the length of the longest <strong>strictly increasing </strong></em><span data-keyword="subsequence-array"><em><strong>subsequence</strong></em></span>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [10,9,2,5,3,7,101,18]
<strong>Output:</strong> 4
<strong>Explanation:</strong> The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,1,0,3,2,3]
<strong>Output:</strong> 4
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [7,7,7,7,7,7,7]
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 2500</code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>

<p>&nbsp;</p>
<p><b>Follow up:</b>&nbsp;Can you come up with an algorithm that runs in&nbsp;<code>O(n log(n))</code> time complexity?</p>

**/
/**
 * [10,9,2,5,3,7,101,18]
**/

// https://leetcode.com/problems/longest-increasing-subsequence/solutions/1366486/golang-binary-search-time-o-nlogn-space-o-n/?orderBy=most_votes&languageTags=golang
func lengthOfLIS(nums []int) int {
	tails := []int{}

	for _, num := range nums {
		idx := binarySearchLengthOfLIS(tails, num)

		if len(tails) == idx {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}

func binarySearchLengthOfLIS(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
