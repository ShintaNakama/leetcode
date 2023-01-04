package main

import "math"

/**
 * <p>You are given an integer array <code>prices</code> where <code>prices[i]</code> is the price of a given stock on the <code>i<sup>th</sup></code> day.</p>

<p>On each day, you may decide to buy and/or sell the stock. You can only hold <strong>at most one</strong> share of the stock at any time. However, you can buy it then immediately sell it on the <strong>same day</strong>.</p>

<p>Find and return <em>the <strong>maximum</strong> profit you can achieve</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> prices = [7,1,5,3,6,4]
<strong>Output:</strong> 7
<strong>Explanation:</strong> Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> prices = [1,2,3,4,5]
<strong>Output:</strong> 4
<strong>Explanation:</strong> Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> prices = [7,6,4,3,1]
<strong>Output:</strong> 0
<strong>Explanation:</strong> There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= prices.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= prices[i] &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * [7,1,5,3,6,4]
**/
//func maxProfit2(prices []int) int {
//	subs := make([]int, 0, len(prices))
//	min := math.MinInt64
//	min = prices[0]
//	for _, price := range prices[1:] {
//		if price < min {
//			min = price
//		}
//		s := price - min
//		if s > 0 {
//			subs = append(subs, s)
//			min = price
//		}
//	}
//
//	var res int
//	for _, sub := range subs {
//		res += sub
//	}
//
//	return res
//}

func maxProfit2(prices []int) int {
	min, sub := math.MinInt64, 0
	min = prices[0]
	for _, price := range prices[1:] {
		if price < min {
			min = price
		}
		s := price - min
		if s > 0 {
			sub += s
			min = price
		}
	}

	return sub
}
