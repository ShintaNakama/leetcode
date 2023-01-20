package main

import (
	"container/heap"
)

/**
 * <p>Given an integer array <code>nums</code> and an integer <code>k</code>, return <em>the</em> <code>k</code> <em>most frequent elements</em>. You may return the answer in <strong>any order</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> nums = [1,1,1,2,2,3], k = 2
<strong>Output:</strong> [1,2]
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> nums = [1], k = 1
<strong>Output:</strong> [1]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li><code>k</code> is in the range <code>[1, the number of unique elements in the array]</code>.</li>
	<li>It is <strong>guaranteed</strong> that the answer is <strong>unique</strong>.</li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow up:</strong> Your algorithm&#39;s time complexity must be better than <code>O(n log n)</code>, where n is the array&#39;s size.</p>

**/
/**
 * [1,1,1,2,2,3]
2
**/
func topKFrequent(nums []int, k int) []int {
	// numsをuniquなmapに変換.valueは要素の出現回数
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	//fmt.Println(m)

	q := priorityQueue{}
	heap.Init(&q)

	// heapにPushしていく
	for k, v := range m {
		heap.Push(&q, element{
			num:   k,
			count: v,
		})
	}
	//fmt.Println(q)

	res := make([]int, k)
	for i := k - 1; i >= 0; {
		// Lenがkより大きい場合はheap.Popだけ
		if q.Len() > k {
			_ = heap.Pop(&q)
			continue
		}

		if q.Len() < 1 {
			v := q[0]
			res[i] = v.num
			break
		}

		e := heap.Pop(&q)

		v := e.(element)

		res[i] = v.num

		i--
	}

	return res
}

type element struct {
	num   int
	count int
}

type priorityQueue []element

func (q priorityQueue) Len() int { return len(q) }

func (q priorityQueue) Less(i, j int) bool { return q[i].count < q[j].count }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(element))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
