package main

/**
 * <p>There is an integer array <code>nums</code> sorted in ascending order (with <strong>distinct</strong> values).</p>

<p>Prior to being passed to your function, <code>nums</code> is <strong>possibly rotated</strong> at an unknown pivot index <code>k</code> (<code>1 &lt;= k &lt; nums.length</code>) such that the resulting array is <code>[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]</code> (<strong>0-indexed</strong>). For example, <code>[0,1,2,4,5,6,7]</code> might be rotated at pivot index <code>3</code> and become <code>[4,5,6,7,0,1,2]</code>.</p>

<p>Given the array <code>nums</code> <strong>after</strong> the possible rotation and an integer <code>target</code>, return <em>the index of </em><code>target</code><em> if it is in </em><code>nums</code><em>, or </em><code>-1</code><em> if it is not in </em><code>nums</code>.</p>

<p>You must write an algorithm with <code>O(log n)</code> runtime complexity.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> nums = [4,5,6,7,0,1,2], target = 0
<strong>Output:</strong> 4
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> nums = [4,5,6,7,0,1,2], target = 3
<strong>Output:</strong> -1
</pre><p><strong class="example">Example 3:</strong></p>
<pre><strong>Input:</strong> nums = [1], target = 0
<strong>Output:</strong> -1
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5000</code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li>All values of <code>nums</code> are <strong>unique</strong>.</li>
	<li><code>nums</code> is an ascending array that is possibly rotated.</li>
	<li><code>-10<sup>4</sup> &lt;= target &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * [4,5,6,7,0,1,2]
0
**/

/**
二分探索を使用して、配列内の最小値（ピボット）を見つけます。ピボットは、配列の回転後に最小値が移動した位置です。最小値は元の配列の先頭
要素であるため、回転後の配列はその最小値を中心にソートされています。
ピボットを見つけたら、通常の二分探索を実行しますが、実際のインデックスをピボットを考慮したものに調整します。これにより、回転を無視して
配列内の目的の値を見つけることができます。
**/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return nums[0]
	}

	// pivotの値を特定する
	// 最終的にleftの値がpivot(回転数)となる
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		// midが右端より大きい場合、最小値が右側にあると判断する
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// midが右端より小さい場合、最小値が左側にあると判断する(次の検索範囲はleft < mid)
			right = mid
		}
	}

	// pivotを考慮してtargetを検索する
	// leftとrightにpivotを加算しておくと通常の2分探索で処理できる
	pivot := left
	//fmt.Println(pivot)

	left = pivot
	right = len(nums) - 1 + pivot

	// targetの値を検索するのでleft<=right
	for left <= right {
		mid := left + (right-left)/2
		// midをlen(nums)で剰余すると本来のmidのindexがわかる
		realmid := mid % len(nums)
		//fmt.Println(mid, realmid)
		// あとは通常の2分探索
		// ポイントは回転された後(引数のnums)の順は考えなくてよくて、昇順でsortされている時と同じ考えで良い
		if nums[realmid] > target {
			right = mid - 1
		} else if nums[realmid] < target {
			left = mid + 1
		} else {
			return realmid
		}
	}

	return -1
}
