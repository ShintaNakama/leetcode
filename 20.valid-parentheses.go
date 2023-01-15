package main

/**
 * <p>Given a string <code>s</code> containing just the characters <code>&#39;(&#39;</code>, <code>&#39;)&#39;</code>, <code>&#39;{&#39;</code>, <code>&#39;}&#39;</code>, <code>&#39;[&#39;</code> and <code>&#39;]&#39;</code>, determine if the input string is valid.</p>

<p>An input string is valid if:</p>

<ol>
	<li>Open brackets must be closed by the same type of brackets.</li>
	<li>Open brackets must be closed in the correct order.</li>
	<li>Every close bracket has a corresponding open bracket of the same type.</li>
</ol>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;()&quot;
<strong>Output:</strong> true
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;()[]{}&quot;
<strong>Output:</strong> true
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;(]&quot;
<strong>Output:</strong> false
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>4</sup></code></li>
	<li><code>s</code> consists of parentheses only <code>&#39;()[]{}&#39;</code>.</li>
</ul>

**/
/**
 * "()"
**/

// 自分の解答
func isValid(s string) bool {
	// lenが2未満は即false
	if len(s) < 2 {
		return false
	}

	// mapで対応表を作成
	m := map[rune]rune{
		40:  41,  // (:)
		91:  93,  // [:]
		123: 125, // {:}
	}

	// strsでカッコを開く側のlistを定義. LIFO(Last-In First-Out: 最初に入れたものを最後に出す)
	strs := make([]rune, 0, 10000)
	for _, c := range s {
		switch c {
		// カッコ開く側はリストappend
		case 40, 91, 123:
			strs = append(strs, c)
		default:
			// カッコ閉じる側
			// この時点でカッコ開く側がlistになければfalse
			if len(strs) == 0 {
				return false
			}
			// listの最後の要素と対になる閉じカッコでなければfalse
			if m[strs[len(strs)-1]] != c {
				return false
			}
			// listの最後の要素は使用されたので削除
			strs = strs[0 : len(strs)-1]
		}
	}

	// 最終的にlistのlenが0になればtrue
	if len(strs) == 0 {
		return true
	}

	return false
}
