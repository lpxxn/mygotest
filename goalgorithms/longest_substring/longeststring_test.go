package longest_substring

import (
	"fmt"
	"testing"
)

func TestLongestStr(t *testing.T) {
	str1 := "abcabcbb"
	lenStr, str1 := longest(str1)
	t.Log(str1)
	if lenStr != 3 {
		t.Error()
	}

	str2 := "abxyzpfabcdbea"
	lenStr, str1 = longest(str2)
	t.Log(str1)
	if lenStr != 9 {
		t.Error()
	}

	lenStr2, str2 := testLongest(str1)
	if lenStr != lenStr2 {
		t.Fatal()
	}
	if str1 != str2 {
		t.Fatal()
	}
	t.Log("-----------")
	str1 = "bbbbb"
	lenStr2, str2 = testLongest(str1)
	t.Log(str2)
	if lenStr2 != 1 {
		t.Error()
	}

	//str1 := "abxyzpfabcdbea"
	//t.Log(longestStr(str1))
	//str1 = "abcabcbb"
	//if longestStr(str1) != 3 {
	//	t.Error()
	//}
	//str1 = "bbbbb"
	//if longestStr(str1) != 1 {
	//	t.Error()
	//}
	//
	//str1 = ""
	//if longestStr(str1) != 0 {
	//	t.Error()
	//}
	//str1 = "a"
	//if longestStr(str1) != 1 {
	//	t.Error()
	//}
	//str1 = " "
	//if longestStr(str1) != 1 {
	//	t.Error()
	//}
}

func testLongest(s string) (int, string) {
	m := make(map[byte]int)
	strLen := 0
	str := ""
	for l, r := 0, 0; r < len(s); r++ {
		fmt.Println(m[s[r]])
		//if idx, ok := m[s[r]]; ok {
		l = max(l, m[s[r]])
		//}
		m[s[r]] = r + 1
		strLen = max(strLen, r-l+1)
		if len(str) < strLen {
			str = s[l : r+1]
		}
	}
	return strLen, str
}

func longest(s string) (int, string) {
	m := make(map[byte]int)
	strLen := 0
	str := ""
	for l, r := 0, 0; r < len(s); r++ {
		if idx, ok := m[s[r]]; ok {
			l = max(l, idx)
		}
		m[s[r]] = r + 1
		strLen = max(strLen, r-l+1)
		if len(str) < strLen {
			str = s[l : r+1]
		}
	}
	return strLen, str
}

// 12ms
func longestStr(s string) int {
	m := make(map[byte]int)
	strLen := 0
	for l, r := 0, 0; r < len(s); r++ {
		//if index, ok := m[s[r]]; ok {
		l = max(l, m[s[r]])
		//}
		// 比如 aaaa 如果不加1， strLen 为2
		// Next substring will start after the last occurrence of current character to avoid its repetition.
		// 比如 aaa 上面的 l 就会变成下一个index
		m[s[r]] = r + 1
		strLen = max(strLen, r-l+1)
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
