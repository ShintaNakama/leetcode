package main

import (
	"sort"
	"strings"
)

/**
 * <p>Given an array of strings <code>strs</code>, group <strong>the anagrams</strong> together. You can return the answer in <strong>any order</strong>.</p>

<p>An <strong>Anagram</strong> is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> strs = ["eat","tea","tan","ate","nat","bat"]
<strong>Output:</strong> [["bat"],["nat","tan"],["ate","eat","tea"]]
</pre><p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> strs = [""]
<strong>Output:</strong> [[""]]
</pre><p><strong class="example">Example 3:</strong></p>
<pre><strong>Input:</strong> strs = ["a"]
<strong>Output:</strong> [["a"]]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= strs.length &lt;= 10<sup>4</sup></code></li>
	<li><code>0 &lt;= strs[i].length &lt;= 100</code></li>
	<li><code>strs[i]</code> consists of lowercase English letters.</li>
</ul>

**/
/**
 * ["eat","tea","tan","ate","nat","bat"]
**/

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))
	for _, str := range strs {
		sl := strings.Split(str, "")
		sort.Strings(sl)
		s := strings.Join(sl, ",")
		//fmt.Println(s)
		m[s] = append(m[s], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// sortを使わないやり方もある
// https://leetcode.com/problems/group-anagrams/solutions/434269/golang-17-lines-of-easy-to-understand-solution-without-using-sort/?orderBy=most_votes&languageTags=golang
//func groupAnagrams(strs []string) [][]string {
//	mp := map[[26]int][]string{}
//	for _, s := range strs {
//		k := [26]int{}
//		for i := 0; i < len(s); i++ {
//			k[s[i]-'a'] += 1
//		}
//		mp[k] = append(mp[k], s)
//	}
//	res := [][]string{}
//	for _, v := range mp {
//		res = append(res, v)
//	}
//	return res
//}
