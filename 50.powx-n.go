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

/*
処理内容

引数nが0の場合、1を返します。
引数nが負の場合、xの逆数を計算し、nを正の整数に変換します。
変数resultに1を代入します。
引数nが0になるまで以下を繰り返します。
nの最下位ビットが1の場合、resultにxを掛けます。
xにxを掛け、nを1ビット右シフトします。
resultを返します。
*/

func myPow(x float64, n int) float64 {
	// 引数nが0の場合、1を返します。
	if n == 0 {
		return 1
	}
	// 引数nが負の場合、xの逆数を計算し、nを正の整数に変換します。
	if n < 0 {
		x = 1 / x
		n = -n
		fmt.Println("n < 0")
		fmt.Println("x", x)
		fmt.Println("n", n)
	}
	// 変数resultに1を代入します。
	result := 1.0
	// 引数nが0になるまで以下を繰り返します。
	for n > 0 {
		// nの最下位ビットが1の場合、resultにxを掛けます。
		// ビット演算子のANDビット演算
		fmt.Println("n & 1", n&1)
		if n&1 == 1 {
			result *= x
			fmt.Println("result", result)
		}
		// xにxを掛け、nを1ビット右シフトします。
		x *= x
		fmt.Println("x", x)
		// 複合代入演算子(ライトシフトして代入)
		n >>= 1
		fmt.Println("n", n)
	}
	return result
}
