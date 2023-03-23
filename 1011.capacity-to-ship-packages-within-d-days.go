package main

/**
 * <p>A conveyor belt has packages that must be shipped from one port to another within <code>days</code> days.</p>

<p>The <code>i<sup>th</sup></code> package on the conveyor belt has a weight of <code>weights[i]</code>. Each day, we load the ship with packages on the conveyor belt (in the order given by <code>weights</code>). We may not load more weight than the maximum weight capacity of the ship.</p>

<p>Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within <code>days</code> days.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> weights = [1,2,3,4,5,6,7,8,9,10], days = 5
<strong>Output:</strong> 15
<strong>Explanation:</strong> A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10

Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> weights = [3,2,2,4,1,4], days = 3
<strong>Output:</strong> 6
<strong>Explanation:</strong> A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
1st day: 3, 2
2nd day: 2, 4
3rd day: 1, 4
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> weights = [1,2,3,1,1], days = 4
<strong>Output:</strong> 3
<strong>Explanation:</strong>
1st day: 1
2nd day: 2
3rd day: 3
4th day: 1, 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= days &lt;= weights.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>1 &lt;= weights[i] &lt;= 500</code></li>
</ul>

**/
/**
 * [1,2,3,4,5,6,7,8,9,10]
5
**/

/**
二分探索を使用して、最小の船の重量容量を見つけます。探索範囲は、荷物の最大重量から荷物の総重量までです。
与えられた容量で荷物をすべて出荷できるかどうかを判断する補助関数を作成します。
補助関数を使用して、船の容量が十分かどうかをチェックします。もし十分であれば、探索範囲を狭めて、さらに小さい容量を試みます。そうでなければ、探索範囲を拡大して、より大きい容量を試みます。
**/

func shipWithinDays(weights []int, days int) int {
	// 与えられた要領で日数内に出荷できるかを判定する関数
	canShip := func(weights []int, capacity int, days int) bool {
		currentDayWeight := 0
		daysUsed := 1

		for _, weight := range weights {
			// capacityを超過したら次の日
			if currentDayWeight+weight > capacity {
				currentDayWeight = 0
				daysUsed++
			}

			currentDayWeight += weight

			// 出荷に使用した日数がdaysを超えたら期限の日数で出荷できないといこと
			if daysUsed > days {
				return false
			}
		}

		return true
	}

	// 最大積載をleft, 積載の合計をright
	left, right := maxWeights(weights), sumWeights(weights)

	for left < right {
		mid := left + (right-left)/2
		//fmt.Println("left", left)
		//fmt.Println("right", right)
		//fmt.Println("mid", mid)

		// 期限の日数内に出荷できれば検索範囲を狭めて、出荷できない場合は範囲を広げる
		if canShip(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func maxWeights(weights []int) int {
	maxWeight := 0
	for _, weight := range weights {
		if weight > maxWeight {
			maxWeight = weight
		}
	}
	return maxWeight
}

func sumWeights(weights []int) int {
	total := 0
	for _, weight := range weights {
		total += weight
	}
	return total
}
