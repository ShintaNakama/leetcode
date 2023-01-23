package main

import (
	"container/heap"
)

/**
 * <p>You are given two integer arrays <code>nums1</code> and <code>nums2</code> sorted in <strong>ascending order</strong> and an integer <code>k</code>.</p>

<p>Define a pair <code>(u, v)</code> which consists of one element from the first array and one element from the second array.</p>

<p>Return <em>the</em> <code>k</code> <em>pairs</em> <code>(u<sub>1</sub>, v<sub>1</sub>), (u<sub>2</sub>, v<sub>2</sub>), ..., (u<sub>k</sub>, v<sub>k</sub>)</code> <em>with the smallest sums</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,7,11], nums2 = [2,4,6], k = 3
<strong>Output:</strong> [[1,2],[1,4],[1,6]]
<strong>Explanation:</strong> The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,1,2], nums2 = [1,2,3], k = 2
<strong>Output:</strong> [[1,1],[1,1]]
<strong>Explanation:</strong> The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,2], nums2 = [3], k = 3
<strong>Output:</strong> [[1,3],[2,3]]
<strong>Explanation:</strong> All possible pairs are returned from the sequence: [1,3],[2,3]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length, nums2.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums1[i], nums2[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>nums1</code> and <code>nums2</code> both are sorted in <strong>ascending order</strong>.</li>
	<li><code>1 &lt;= k &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * [1,7,11]
[2,4,6]
3
**/
type Element struct {
	x   int
	y   int
	sum int
}

type PriorityQueue []Element

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].sum < q[j].sum }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *PriorityQueue) Push(x interface{}) {
	*q = append(*q, x.(Element))
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	q := PriorityQueue{}
	heap.Init(&q)

	// nums1の各要素とnums2[0]だけでheapを作る
	for i := range nums1 {
		heap.Push(&q, Element{
			x:   i,
			y:   0,
			sum: nums1[i] + nums2[0],
		})
	}

	//fmt.Println(q)

	res := make([][]int, 0, k)

	// 以下のループでheapの最小要素を取り出しつつnums1[i]とnums2[i+1]の要素との合計もheapにPushしていく
	// つまり前段のループで追加した要素よりもnums1[i]とnums2[i+1]の合計の方が小さい場合はnums1[i]とnums2[i+1]がPopされる
	for k > 0 && q.Len() > 0 {
		// 最小合計値を取り出しappend
		v := heap.Pop(&q)
		e := v.(Element)
		//fmt.Println(e)
		res = append(res, []int{nums1[e.x], nums2[e.y]})

		if len(res) == k {
			break
		}

		//fmt.Println(len(res))
		if e.y+1 < len(nums2) {
			heap.Push(&q, Element{
				x:   e.x,
				y:   e.y + 1,
				sum: nums1[e.x] + (nums2[e.y+1]),
			})
		}
	}

	return res
}
