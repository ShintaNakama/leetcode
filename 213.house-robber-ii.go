package main

/**
 * <p>You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are <strong>arranged in a circle.</strong> That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and&nbsp;<b>it will automatically contact the police if two adjacent houses were broken into on the same night</b>.</p>

<p>Given an integer array <code>nums</code> representing the amount of money of each house, return <em>the maximum amount of money you can rob tonight <strong>without alerting the police</strong></em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,3,2]
<strong>Output:</strong> 3
<strong>Explanation:</strong> You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,3,1]
<strong>Output:</strong> 4
<strong>Explanation:</strong> Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,3]
<strong>Output:</strong> 3
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 1000</code></li>
</ul>

**/
/**
 * [2,3,2]
**/
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var (
		included int
		excluded int
	)

	for i, n := range nums {
		tmp := included

		if i != len(nums)-1 {
			included = excluded + n
		}

		excluded = max(tmp, excluded)
	}

	n1 := max(included, excluded)

	var (
		included2 int
		excluded2 int
	)

	for i := 1; i < len(nums); i++ {
		tmp := included2
		included2 = excluded2 + nums[i]
		excluded2 = max(tmp, excluded2)
	}

	n2 := max(included2, excluded2)

	return max(n1, n2)
}
