package main

/**
 * <p>Given a string <code>s</code>, <em>find the first non-repeating character in it and return its index</em>. If it does not exist, return <code>-1</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> s = "leetcode"
<strong>Output:</strong> 0
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> s = "loveleetcode"
<strong>Output:</strong> 2
</pre><p><strong class="example">Example 3:</strong></p>
<pre><strong>Input:</strong> s = "aabb"
<strong>Output:</strong> -1
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> consists of only lowercase English letters.</li>
</ul>

**/
/**
 * "leetcode"
**/
//func firstUniqChar(s string) int {
//	m := make(map[string]int, len(s))
//	m2 := make(map[string]int, len(s))
//
//	ss := strings.Split(s, "")
//
//	for i, s := range ss {
//		if _, ok := m[s]; ok {
//			m2[s]++
//			continue
//		}
//		m[s] = i
//		m2[s] = 1
//	}
//
//	index := len(s)
//	for k2, v2 := range m2 {
//		if v2 > 1 {
//			continue
//		}
//		if v, ok := m[k2]; ok {
//			if index > v {
//				index = v
//			}
//		}
//	}
//
//	if index == len(s) {
//		return -1
//	}
//
//	return index
//}

// 0ms
//https://leetcode.com/problems/first-unique-character-in-a-string/solutions/1961087/go-solution-in-o-n-time-and-constant-space-complexity-0ms-100-faster/?orderBy=most_votes&languageTags=golang
func firstUniqChar(s string) int {
	// This is a constant space allocation (ie: always 26)
	var counts = make([]int, 26)

	// Insert all the character's count into counts array
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	// Find the first element with count 1
	for i := 0; i < len(s); i++ {
		if counts[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// mapを使ったシンプルなパターン
//https://leetcode.com/problems/first-unique-character-in-a-string/solutions/337965/go-o-n-hashmap-with-explanation/?orderBy=most_votes&languageTags=golang
//func firstUniqChar(s string) int {
//	d := map[byte]int{}
//
//	// Count each character.
//	for i := 0; i < len(s); i++ {
//		d[s[i]]++
//	}
//
//	// Find the first unique character and return.
//	for i := 0; i < len(s); i++ {
//		if d[s[i]] == 1 {
//			return i
//		}
//	}
//
//	// If there's no unique character then return -1.
//	return -1
//}
