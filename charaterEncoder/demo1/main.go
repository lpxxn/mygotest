package main

import (
	"bytes"
	"fmt"
	"time"

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

	const dateTimeFormat = "060102150405"
	updateTime := time.Now()
	buf := new(bytes.Buffer)

	// 餐品的最后修改时间, 格式 YYMMDDhhmmss, ASCII 32 bytes, 空位补 0.
	timeStr := updateTime.Format(dateTimeFormat)
	timeBytes := make([]byte, 0, 32)
	timeBytes = append(timeBytes, []byte(timeStr)...)
	fmt.Println(timeBytes, len(timeBytes))
	// 补足 32 bytes.
	timeBytes = append(timeBytes, make([]byte, 32-len(timeBytes))...)
	fmt.Println(timeBytes, len(timeBytes))
	fmt.Println(string(timeBytes), len(timeBytes))
	buf.Write(timeBytes)
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
