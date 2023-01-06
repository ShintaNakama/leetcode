package main

import "fmt"

/**
 * <p>Implement <a href="http://www.cplusplus.com/reference/valarray/pow/" target="_blank">pow(x, n)</a>, which calculates <code>x</code> raised to the power <code>n</code> (i.e., <code>x<sup>n</sup></code>).</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> x = 2.00000, n = 10
<strong>Output:</strong> 1024.00000
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> x = 2.10000, n = 3
<strong>Output:</strong> 9.26100
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> x = 2.00000, n = -2
<strong>Output:</strong> 0.25000
<strong>Explanation:</strong> 2<sup>-2</sup> = 1/2<sup>2</sup> = 1/4 = 0.25
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-100.0 &lt; x &lt; 100.0</code></li>
	<li><code>-2<sup>31</sup> &lt;= n &lt;= 2<sup>31</sup>-1</code></li>
	<li><code>n</code> is an integer.</li>
	<li><code>-10<sup>4</sup> &lt;= x<sup>n</sup> &lt;= 10<sup>4</sup></code></li>
</ul>

**/
/**
 * 2.00000
10
*/

// https://leetcode.com/problems/powx-n/solutions/1702631/simple-golang-solution-100-faster-o-log-n-recursive/?orderBy=most_votes&languageTags=golang

func myPow(x float64, n int) float64 {
	fmt.Println("nn:", n)
	if n == 0 {
		return 1.0
	} else if n < 0 {
		//fmt.Println(myPow(x, -n))
		return 1.0 / myPow(x, -n)
	}

	// n/2 を引数に再帰的に計算する
	tmp := myPow(x, n/2)
	//fmt.Println("n:", n)
	//fmt.Println("tmp:", tmp)
	//fmt.Println("n%2", n%2)

	// no1だと最終的に32*32が戻り値になる(32=2の5乗)
	if n%2 == 0 {
		return tmp * tmp
	} else {
		return x * tmp * tmp
	}
}
