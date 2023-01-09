package main

import "fmt"

/**
 * <p>Given an integer array <code>nums</code>, move all <code>0</code>&#39;s to the end of it while maintaining the relative order of the non-zero elements.</p>

<p><strong>Note</strong> that you must do this in-place without making a copy of the array.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> nums = [0,1,0,3,12]
<strong>Output:</strong> [1,3,12,0,0]
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> nums = [0]
<strong>Output:</strong> [0]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>-2<sup>31</sup> &lt;= nums[i] &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you minimize the total number of operations done?
**/
/**
 * [0,1,0,3,12]
 */

// https://leetcode.com/problems/move-zeroes/solutions/724397/idiomatic-go-solution/?orderBy=most_votes&languageTags=golang
func moveZeroes(nums []int) {
	// 0でない数値のindex
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		// nums[i]が0でない場合はnums[i]をnums[nonZeroIndex]を入れ替え
		if nums[i] != 0 {
			//fmt.Println(i)
			//fmt.Println(nums[i])
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			//fmt.Println(nums[i])
			nonZeroIndex++
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)
}
