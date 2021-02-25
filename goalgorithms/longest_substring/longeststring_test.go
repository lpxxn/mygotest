package longest_substring

import "testing"

func TestLongestStr(t *testing.T) {
	str := "abxyzpfabcdbea"
	t.Log(longestStr(str))
	str = "abcabcbb"
	if longestStr(str) != 3 {
		t.Error()
	}
	str = "bbbbb"
	if longestStr(str) != 1 {
		t.Error()
	}

	str = ""
	if longestStr(str) != 0 {
		t.Error()
	}
	str = "a"
	if longestStr(str) != 1 {
		t.Error()
	}
	str = " "
	if longestStr(str) != 1 {
		t.Error()
	}
}

// 12ms
func longestStr(s string) int {
	m := make(map[byte]int)
	strLen := 0
	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[s[r]]; ok {
			l = max(l, index)
		}
		strLen = max(strLen, r-l+1)
		// 比如 aaaa 如果不加1， strLen 为2
		// Next substring will start after the last occurrence of current character to avoid its repetition.
		// 比如 aaa 上面的 l 就会变成下一个index
		m[s[r]] = r + 1
	}
	return strLen
}

// 0 ms
func longestSubString(s string) int {
	// ASCII 0 -- 127
	m := [128]int{}
	longest := 0
	left := 0
	for right := 0; right < len(s); right++ {
		// s[right] 是新遇到的字母
		// m 是记录字母对应的位置 + 1是因为，如果有重复的，就直接取下一个字母的index 也就是 +1 的index，
		// 所以 m[s[right]] 就是新字母对应的新位置
		left = max(left, m[s[right]])
		m[s[right]] = right + 1
		longest = max(longest, right-left+1)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
