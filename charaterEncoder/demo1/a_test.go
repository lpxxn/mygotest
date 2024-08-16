package main

import (
	"sort"
	"testing"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func TestChinese(t *testing.T) {
	collator := collate.New(language.SimplifiedChinese)
	data := []string{"张三", "李四", "王五", "李一"}
	sort.Slice(data, func(i, j int) bool {
		return collator.CompareString(data[i], data[j]) < 0
	})
	println(data)
}

func generateCombinations(length int) []string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return generateCombinationsRecursive("", letters, length)
}
 // all length 指定长度的组合
func generateCombinationsRecursive(prefix string, letters string, remaining int) []string {
	if remaining == 0 {
		return []string{prefix}
	}

	var combinations []string
	for i := 0; i < len(letters); i++ {
		newPrefix := prefix + string(letters[i])
		combinations = append(combinations, generateCombinationsRecursive(newPrefix, letters, remaining-1)...)
	}

	return combinations
}

func TestGeneStr(t *testing.T) {
	length := 3 // 你可以更改这个值来生成不同长度的组合
	combinations := generateCombinations(length)

	t.Logf("Total combinations of length %d: %d\n", length, len(combinations))

	// 打印前几个组合作为示例
	maxPrint := 20
	for i := 0; i < maxPrint && i < len(combinations); i++ {
		t.Logf(combinations[i])
	}

	if len(combinations) > maxPrint {
		t.Logf("... and %d more combinations\n", len(combinations)-maxPrint)
	}
}
