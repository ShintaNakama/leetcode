package main

/**
 * <p>Suppose an array of length <code>n</code> sorted in ascending order is <strong>rotated</strong> between <code>1</code> and <code>n</code> times. For example, the array <code>nums = [0,1,2,4,5,6,7]</code> might become:</p>

<ul>
	<li><code>[4,5,6,7,0,1,2]</code> if it was rotated <code>4</code> times.</li>
	<li><code>[0,1,2,4,5,6,7]</code> if it was rotated <code>7</code> times.</li>
</ul>

<p>Notice that <strong>rotating</strong> an array <code>[a[0], a[1], a[2], ..., a[n-1]]</code> 1 time results in the array <code>[a[n-1], a[0], a[1], a[2], ..., a[n-2]]</code>.</p>

<p>Given the sorted rotated array <code>nums</code> of <strong>unique</strong> elements, return <em>the minimum element of this array</em>.</p>

<p>You must write an algorithm that runs in&nbsp;<code>O(log n) time.</code></p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,4,5,1,2]
<strong>Output:</strong> 1
<strong>Explanation:</strong> The original array was [1,2,3,4,5] rotated 3 times.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [4,5,6,7,0,1,2]
<strong>Output:</strong> 0
<strong>Explanation:</strong> The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [11,13,15,17]
<strong>Output:</strong> 11
<strong>Explanation:</strong> The original array was [11,13,15,17] and it was rotated 4 times.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>n == nums.length</code></li>
	<li><code>1 &lt;= n &lt;= 5000</code></li>
	<li><code>-5000 &lt;= nums[i] &lt;= 5000</code></li>
	<li>All the integers of <code>nums</code> are <strong>unique</strong>.</li>
	<li><code>nums</code> is sorted and rotated between <code>1</code> and <code>n</code> times.</li>
</ul>

**/
/**
 * [3,4,5,1,2]
**/

/**
「回転したソートされた配列」で最小要素を求める問題
二分探索を使用することでO(log n)の時間計算量で解決でる。

1. 配列の中央の要素を取得
2. 中央の要素が最小値である場合、中央の要素を返す。
3. 中央の要素が左半分に属する要素よりも小さい場合、最小値は左半分に存在するので、左半分に対して再帰的にこの手順を繰り返す。
4. 中央の要素が右半分に属する要素よりも大きい場合、最小値は右半分に存在するので、右半分に対して再帰的にこの手順を繰り返す。

leftとrightは配列の先頭と末尾を指すポインタであり、whileループの条件はleft < rightとなっています。ループの中で、中央の要素を取得して、その要素が右側の要素よりも大きい場合は、最小値は右半分に存在するため、leftをmid+1に設定します。それ以外の場合は、最小値は左半分に存在するため、rightをmidに設定します。

最終的に、leftとrightが重なると、その要素が配列の最小要素となります。
**/

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		//fmt.Println(mid)
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}

	}

	return nums[left]
}
