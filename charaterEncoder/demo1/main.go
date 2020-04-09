package main

import (
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	encoderStr("红烧东海三鲜包含水潺、带鱼、鲳鱼")
	encoderStr("✨测试餐品")
	encoderStr("🔥猪杂猪肉丸汤粉（大）")
	encoderStr("1")
	encoderStr("0.5")
	encoderStr("<食堂菜单>2018-07-19~2018-07-31[按天重复_11]") // 44大于 40
	encoderStr("2019-03-17 起[按天重复_27]")
}

func encoderStr(s string) {
	v, err := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(s))
	if err != nil {
		panic(err)
	}
	fmt.Println("v: ", v, " len: ", len(v), " u8: ", uint8(len(v)))
	dv, err := simplifiedchinese.GB18030.NewDecoder().Bytes(v)
	if err != nil {
		panic(err)
	}
	fmt.Println("dv: ", string(dv), " len: ", len(dv))
}
