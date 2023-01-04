package main

/**
 * <p>Given two integer arrays <code>nums1</code> and <code>nums2</code>, return <em>an array of their intersection</em>. Each element in the result must be <strong>unique</strong> and you may return the result in <strong>any order</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,2,2,1], nums2 = [2,2]
<strong>Output:</strong> [2]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [4,9,5], nums2 = [9,4,9,8,4]
<strong>Output:</strong> [9,4]
<strong>Explanation:</strong> [4,9] is also accepted.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length, nums2.length &lt;= 1000</code></li>
	<li><code>0 &lt;= nums1[i], nums2[i] &lt;= 1000</code></li>
</ul>

**/
/**
 * [1,2,2,1]
[2,2]
**/
func intersection(nums1 []int, nums2 []int) []int {
	longestLen := len(nums1)
	if len(nums1) < len(nums2) {
		longestLen = len(nums2)
	}

	m := make(map[int]bool, len(nums1))

	for _, n := range nums1 {
		m[n] = true
	}

	res := make([]int, 0, longestLen)
	for _, n := range nums2 {
		v, ok := m[n]
		if ok && v {
			res = append(res, n)
			//delete(m, n)
			m[n] = false
		}
	}

	return res
}

// 最初の回答.acceptした
//func intersection(nums1 []int, nums2 []int) []int {
//	longestLen := len(nums1)
//	if len(nums1) < len(nums2) {
//		longestLen = len(nums2)
//	}
//
//	m := make(map[int]struct{}, len(nums1))
//
//	for _, n := range nums1 {
//		m[n] = struct{}{}
//	}
//
//	resM := make(map[int]struct{}, longestLen)
//
//	for _, n := range nums2 {
//		_, ok := m[n]
//		if ok {
//			resM[n] = struct{}{}
//		}
//	}
//
//	res := make([]int, 0, longestLen)
//	for k := range resM {
//		res = append(res, k)
//	}
//
//	return res
//}
