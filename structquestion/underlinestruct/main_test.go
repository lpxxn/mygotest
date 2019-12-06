package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	Name string `json:"name"`
}
type AList []*A

func TestStructArray(t *testing.T) {
	// "[]"
	a1 := AList{}
	fmt.Printf("%#v \n", a1)
	b1, _ := json.Marshal(a1)
	fmt.Printf("%#v \n", string(b1))

	// "[]"
	a2 := make(AList, 0, 1)
	fmt.Printf("%#v \n", a2)
	b2, _ := json.Marshal(a2)
	fmt.Printf("%#v \n", string(b2))

	// "null"
	var a3 AList
	fmt.Printf("%#v \n", a3)
	b3, _ := json.Marshal(a3)
	fmt.Printf("%#v \n", string(b3))

}

// Color QRCode color
type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

// UnlimitedQRCode
type UnlimitedQRCode struct {
	// 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：
	Scene string `json:"scene"`
	Page  string `json:"page,omitempty"`
	// 二维码的宽度，单位 px，最小 280px，最大 1280px
	Width     int  `json:"width,omitempty"`
	AutoColor bool `json:"auto_color,omitempty"`
	// auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	LineColor Color `json:"line_color,omitempty"`
	// 是否需要透明底色，为 true 时，生成透明底色的小程序
	IsHyaline bool `json:"is_hyaline,omitempty"`
}

func TestCompareStruct(t *testing.T) {
	s1 := &UnlimitedQRCode{
		Scene:     "",
		Page:      "",
		Width:     0,
		AutoColor: false,
		LineColor: Color{},
		IsHyaline: false,
	}
	s2 := &UnlimitedQRCode{
		Scene:     "",
		Page:      "",
		Width:     0,
		AutoColor: false,
		LineColor: Color{},
		IsHyaline: false,
	}

	if s1 == s2 {
		t.Fatal("pointer equal error")
	}

	if *s1 != *s2 {
		t.Fatal("not equal")
	}
	s1.LineColor.B = "aa"

	if *s1 == *s2 {
		t.Fatal("not equal")
	}
	s2.LineColor.B = "aa"
	s1.Scene = "scene=2002?cid=1&sid=1"
	s2.Scene = "scene=2002?cid=1&sid=1"
	if *s1 != *s2 {
		t.Fatal("not equal")
	}
}
