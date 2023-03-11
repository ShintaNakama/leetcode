package main

import "fmt"

/**
 * <p>You are given an integer array <code>coins</code> representing coins of different denominations and an integer <code>amount</code> representing a total amount of money.</p>

<p>Return <em>the fewest number of coins that you need to make up that amount</em>. If that amount of money cannot be made up by any combination of the coins, return <code>-1</code>.</p>

<p>You may assume that you have an infinite number of each kind of coin.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> coins = [1,2,5], amount = 11
<strong>Output:</strong> 3
<strong>Explanation:</strong> 11 = 5 + 5 + 1
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> coins = [2], amount = 3
<strong>Output:</strong> -1
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> coins = [1], amount = 0
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= coins.length &lt;= 12</code></li>
	<li><code>1 &lt;= coins[i] &lt;= 2<sup>31</sup> - 1</code></li>
	<li><code>0 &lt;= amount &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * [1,2,5]
11
**/

// 例coins[1,2,5] amount:11の場合
// このアルゴリズムは、0から11までの金額で必要な最小のコインの枚数を計算
// dp配列の最後がamount:11になる最小のコイン枚数となる
func coinChange(coins []int, amount int) int {
	// amount+1を初期値としたdp配列を作ることで
	// 0..amountまでの最小の枚数を記憶しておけるようにする
	maxAmount := amount + 1
	dp := make([]int, maxAmount)
	for i := 1; i < maxAmount; i++ {
		dp[i] = maxAmount
	}

	fmt.Println(dp)

	// 0は0なので、1..amountまでの金額を何枚のコインで再現できるか計算する
	// そのため1..amount+1のループでさらに内部でcoinsでループ
	for i := 1; i < maxAmount; i++ {
		for _, coin := range coins {
			// iがコインで再現したい金額
			// iがcoinより大きいとそのcoinでは再現できないのでskip
			if coin <= i {
				// ここでi-coinはiの前の金額。
				// dp[i-coin]はこの前の金額に到達するために必要な最小のコインの枚数。
				// dp[i-coin]+1は現在のコインを1枚追加することによって、金額iに到達するために必要な最小コイン数を表す。
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	fmt.Println(dp)

	if dp[amount] == maxAmount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
