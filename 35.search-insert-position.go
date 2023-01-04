package main

/**
 * <p>Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.</p>

<p>You must&nbsp;write an algorithm with&nbsp;<code>O(log n)</code> runtime complexity.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,5,6], target = 5
<strong>Output:</strong> 2
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,5,6], target = 2
<strong>Output:</strong> 1
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,5,6], target = 7
<strong>Output:</strong> 4
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li><code>nums</code> contains <strong>distinct</strong> values sorted in <strong>ascending</strong> order.</li>
	<li><code>-10<sup>4</sup> &lt;= target &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * [1,3,5,6]
5
**/

// https://leetcode.com/problems/search-insert-position/solutions/519363/python-js-go-c-o-log-n-sol-by-binary-search-with-hint/?orderBy=most_votes&languageTags=golang
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// Classical binary search
	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid - 1

		} else if nums[mid] < target {
			left = mid + 1

		} else {
			// Target exists in nums, so the index of repeated element is the insert position
			return mid
		}

	}

	// Target doesn't exist in nums, so the index of first largest element is the insert position
	return left
}
