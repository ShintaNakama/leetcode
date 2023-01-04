package main

import (
	"strings"
)

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

// 自分の回答
func isValid(s string) bool {
	ss := strings.Split(s, "")

	cc := make([]string, 0, 10)

	for i, c := range ss {
		if i == 0 {
			cc = append(cc, c)
			continue
		}

		switch c {
		case ")":
			if len(cc) != 0 && cc[len(cc)-1] == "(" {
				cc = cc[0 : len(cc)-1]
				continue
			} else {
				return false
			}
		case "}":
			if len(cc) != 0 && cc[len(cc)-1] == "{" {
				cc = cc[0 : len(cc)-1]
				continue
			} else {
				return false
			}
		case "]":
			if len(cc) != 0 && cc[len(cc)-1] == "[" {
				cc = cc[0 : len(cc)-1]
				continue
			} else {
				return false
			}
		case "(":
			cc = append(cc, c)
		case "{":
			cc = append(cc, c)
		case "[":
			cc = append(cc, c)
		}
	}

	return len(cc) == 0
}

// https://leetcode.com/problems/valid-parentheses/solutions/272483/golang-solution-0ms-2mb-100/?orderBy=most_votes&languageTags=golang
//func isValid(s string) bool {
//    if s == "" {
//        return true
//    }
//
//    tr := map[rune]rune{
//        '}': '{',
//        ')': '(',
//        ']': '[',
//    }
//
//    stack := make([]rune, 0, 1)
//
//    for _, ch := range s {
//        switch ch {
//        case 40,123,91:
//            stack = append(stack, ch)
//        case 41,125,93:
//            if len(stack) == 0 || stack[len(stack)-1] != tr[ch] {
//                return false
//            } else {
//                stack = stack[:len(stack)-1]
//            }
//        }
//    }
//
//    return len(stack) == 0
//}
