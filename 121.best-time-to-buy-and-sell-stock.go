package main

import "math"

/**
 * <p>You are given an array <code>prices</code> where <code>prices[i]</code> is the price of a given stock on the <code>i<sup>th</sup></code> day.</p>

<p>You want to maximize your profit by choosing a <strong>single day</strong> to buy one stock and choosing a <strong>different day in the future</strong> to sell that stock.</p>

<p>Return <em>the maximum profit you can achieve from this transaction</em>. If you cannot achieve any profit, return <code>0</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> prices = [7,1,5,3,6,4]
<strong>Output:</strong> 5
<strong>Explanation:</strong> Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> prices = [7,6,4,3,1]
<strong>Output:</strong> 0
<strong>Explanation:</strong> In this case, no transactions are done and the max profit = 0.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= prices.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= prices[i] &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * [7,1,5,3,6,4]
**/
//// 自分の回答
//func maxProfit(prices []int) int {
//	min, sub := math.MinInt64, 0
//	min = prices[0]
//	for _, price := range prices[1:] {
//		if price < min {
//			min = price
//		}
//		s := price - min
//		if s > sub {
//			sub = s
//		}
//
//	}
//
//	return sub
//}

// 自分の回答2
func maxProfit(prices []int) int {
	min, sub := math.MinInt64, 0
	min = prices[0]
	for _, price := range prices[1:] {
		if price < min {
			min = price
		}
		s := price - min
		if s > sub {
			sub = s
		}

	}

	return sub
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/802390/python-go-js-c-o-n-by-dp-w-visualization/?orderBy=most_votes&languageTags=golang
