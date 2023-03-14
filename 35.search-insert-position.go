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

// 模範回答
// https://leetcode.com/problems/search-insert-position/solutions/1787521/go-concise-100/?orderBy=most_votes&languageTags=golang
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		herf := ((right - left) / 2) + left
		//fmt.Println(herf)
		v := nums[herf]
		switch {
		case v < target:
			left = herf + 1
		case v > target:
			right = herf - 1
		default:
			return herf
		}
	}

	return left
}

// 自分の回答
//func searchInsert(nums []int, target int) int {
//	for {
//		if len(nums) <= 2 {
//			if len(nums) == 2 {
//				if nums[0] < target && nums[1] > target {
//					return 1
//				} else if nums[1] < target {
//					return 2
//				}
//			}
//			if nums[0] > target {
//				return 0
//			} else if nums[0] < target {
//				return 1
//			}
//		}
//		herf := len(nums) / 2
//
//		n := nums[herf]
//
//		if n == target {
//			return herf
//		}
//
//		if target >= n {
//			right := searchInsert(nums[herf:], target)
//			return right + herf
//		} else {
//			left := searchInsert(nums[0:herf], target)
//			return left
//		}
//	}
//}
//
