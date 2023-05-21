package main

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"sort"
	"testing"
)

func TestChinese(t *testing.T) {
	collator := collate.New(language.SimplifiedChinese)
	data := []string{"张三", "李四", "王五", "李一"}
	sort.Slice(data, func(i, j int) bool {
		return collator.CompareString(data[i], data[j]) < 0
	})
	println(data)
}
