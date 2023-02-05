package main

/**
 * <p>A <strong>transformation sequence</strong> from word <code>beginWord</code> to word <code>endWord</code> using a dictionary <code>wordList</code> is a sequence of words <code>beginWord -&gt; s<sub>1</sub> -&gt; s<sub>2</sub> -&gt; ... -&gt; s<sub>k</sub></code> such that:</p>

<ul>
	<li>Every adjacent pair of words differs by a single letter.</li>
	<li>Every <code>s<sub>i</sub></code> for <code>1 &lt;= i &lt;= k</code> is in <code>wordList</code>. Note that <code>beginWord</code> does not need to be in <code>wordList</code>.</li>
	<li><code>s<sub>k</sub> == endWord</code></li>
</ul>

<p>Given two words, <code>beginWord</code> and <code>endWord</code>, and a dictionary <code>wordList</code>, return <em>the <strong>number of words</strong> in the <strong>shortest transformation sequence</strong> from</em> <code>beginWord</code> <em>to</em> <code>endWord</code><em>, or </em><code>0</code><em> if no such sequence exists.</em></p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> beginWord = &quot;hit&quot;, endWord = &quot;cog&quot;, wordList = [&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;lot&quot;,&quot;log&quot;,&quot;cog&quot;]
<strong>Output:</strong> 5
<strong>Explanation:</strong> One shortest transformation sequence is &quot;hit&quot; -&gt; &quot;hot&quot; -&gt; &quot;dot&quot; -&gt; &quot;dog&quot; -&gt; cog&quot;, which is 5 words long.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> beginWord = &quot;hit&quot;, endWord = &quot;cog&quot;, wordList = [&quot;hot&quot;,&quot;dot&quot;,&quot;dog&quot;,&quot;lot&quot;,&quot;log&quot;]
<strong>Output:</strong> 0
<strong>Explanation:</strong> The endWord &quot;cog&quot; is not in wordList, therefore there is no valid transformation sequence.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= beginWord.length &lt;= 10</code></li>
	<li><code>endWord.length == beginWord.length</code></li>
	<li><code>1 &lt;= wordList.length &lt;= 5000</code></li>
	<li><code>wordList[i].length == beginWord.length</code></li>
	<li><code>beginWord</code>, <code>endWord</code>, and <code>wordList[i]</code> consist of lowercase English letters.</li>
	<li><code>beginWord != endWord</code></li>
	<li>All the words in <code>wordList</code> are <strong>unique</strong>.</li>
</ul>

**/
/**
 * "hit"
"cog"
["hot","dot","dog","lot","log","cog"]
**/

// 参考:https://leetcode.com/problems/word-ladder/solutions/3136332/go-bfs/?orderBy=most_votes&languageTags=golang
type W struct {
	word string
	step int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// wordListで探索済みかどうかを記録するmapを生成
	// 値はfalseとし探索したらtrueにする
	m := map[string]bool{}

	for _, word := range wordList {
		m[word] = false
	}

	// queueを生成
	q := []W{
		{
			word: beginWord,
			step: 1,
		},
	}

	// BFS(幅優先探索)
	// qに要素が存在する限りループ
	for len(q) > 0 {
		// qから先頭の要素を取り出す
		w := q[0]
		q = q[1:]

		// wordMapに存在しかつ探索済み(値がtrue)の場合continue
		// つまりstepは進まない
		if v, ok := m[w.word]; ok && v {
			continue
		}

		// 訪問したwordはtrueにする
		m[w.word] = true

		// endWordとqから取り出したwordが一致したらstepを返し終了
		if w.word == endWord {
			return w.step
		}

		// wordListをループしマッチした文字列はappend
		// この時qから取り出したwordの要素がもつstepに+1する
		for _, word := range wordList {
			if diffStr(w.word, word) {
				q = append(q, W{
					word: word,
					step: w.step + 1,
				})
			}
		}
	}

	return 0
}

func diffStr(s1, s2 string) bool {
	var diff int
	for i := range s1 {
		if s1[i] != s2[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}

	return diff == 1
}
