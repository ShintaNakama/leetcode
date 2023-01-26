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
	longNums := nums1
	shortNums := nums2
	if len(nums1) < len(nums2) {
		longNums = nums2
		shortNums = nums1
	}

	m := make(map[int]struct{}, len(longNums))
	for _, n := range longNums {
		m[n] = struct{}{}
	}

	m2 := make(map[int]struct{}, len(shortNums))
	for _, n := range shortNums {
		if _, ok := m[n]; ok {
			m2[n] = struct{}{}
		}
	}

	res := make([]int, 0, len(m2))
	for k := range m2 {
		res = append(res, k)
	}

	return res
}

// もう少し効率の良いやり方がある
// https://leetcode.com/problems/intersection-of-two-arrays/solutions/405483/go-solution/?orderBy=most_votes&languageTags=golang
// func intersection(nums1 []int, nums2 []int) []int {
//     var count = map[int]bool{}
//     var result = []int{}
//     for _,num := range nums1{
//         count[num]=true
//     }
//     for _, num := range nums2{
//         if(count[num]==true){
//             result=append(result,num)
//             count[num]=false
//         }
//     }
//     return result
// }
