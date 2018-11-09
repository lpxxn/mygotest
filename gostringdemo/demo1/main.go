package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Atest struct {
	Name string
	Age  int
}

func (a Atest) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}

func main() {
	fmt.Println(encode("1386"))
	a := Atest{Name: "li", Age: 18}
	s := fmt.Sprintf("%v", a)
	fmt.Println(s)

	s1 := "1234567"
	fmt.Println(s1[2:5])

	h := md5.New()
	io.WriteString(h, "http:@@@@@@www.baidu.com")
	m1 := h.Sum(nil)
	fmt.Printf("%x\n", m1)

	b641 := base64.StdEncoding.EncodeToString(m1)
	fmt.Println(b641)

	rev := encode("http://www.baidu.com")
	fmt.Println(rev)

}

type CardType struct {
	CoverImg  string
	BackImg   string
	CardParam CardTypeParam
}

type CardTypeParam struct {
	No string
	Pattern []int
}

const (
	BACK_IMG_PATTERN        = "/card/b/%s/%s"
	COVER_IMG_PATTERN       = "/card/c/%s/%s"
	ENCODE_PATTERN          = "%s/mvalue/%s"
	BACK_EMPTY_IMG_PATTERN  = "/card/back/%s"
	COVER_EMPTY_IMG_PATTERN = "/card/cover/%s"
)

/**
 * 返回经过pattern模式的解析之后.每一个具体的数字是多少
 *
 * @param cardNumber
 * @param pattern
 * @return
 */
func (*CardType) applyCardNumberPattern(cardNumber string, pattern []int) []string {
	totalSize := len(pattern)
	applyResult := make([]string, totalSize)
	lastIndex := 0
	for i := 0; i < totalSize; i++ {
		item := pattern[i]
		endIndex := lastIndex + item

		applyResult[i] = cardNumber[lastIndex:endIndex]
		lastIndex = endIndex
	}
	return applyResult
}

func (c *CardType) BuildCoverImageUrl(cardNumber string) string {
	if len(cardNumber) == 0 {
		return fmt.Sprintf(COVER_EMPTY_IMG_PATTERN, c.CoverImg)
	}
	return c.buildNonEmptyCoverImageUrl(cardNumber)
}

func (c *CardType) buildBackImageUrl(cardNumber string) string {
	if len(cardNumber) == 0 {
		return fmt.Sprintf(BACK_EMPTY_IMG_PATTERN, c.CoverImg)
	}
	return c.buildNonEmptyBackImageUrl(cardNumber)
}

func (c *CardType) buildNonEmptyBackImageUrl(cardNumber string) string {

	cardNumberInPattern := c.applyCardNumberPattern(cardNumber, c.CardParam.Pattern)
	/**
	 *构造卡片的图片url的时候使用的是"-"分割卡号.因为每一部分可能会显示为不同的格式.这样在模板里就可以直接配变量了.
	 */
	joinedCardNumber := strings.Join(cardNumberInPattern, "-")
	base := fmt.Sprintf(BACK_IMG_PATTERN, joinedCardNumber, c.BackImg)
	check := encode(base)
	return fmt.Sprintf(ENCODE_PATTERN, base, check)
}

func (c *CardType) buildNonEmptyCoverImageUrl(cardNumber string) string {
	cardNumberInPattern := c.applyCardNumberPattern(cardNumber, c.CardParam.Pattern)

	/**
	 *构造卡片的图片url的时候使用的是"-"分割卡号.因为每一部分可能会显示为不同的格式.这样在模板里就可以直接配变量了.
	 */
	joinedCardNumber := strings.Join(cardNumberInPattern, "-")
	base := fmt.Sprintf(COVER_IMG_PATTERN, joinedCardNumber, c.CoverImg)
	check := encode(base)
	return fmt.Sprintf(ENCODE_PATTERN, base, check)
}

func encode(uri string) string {
	repStr := strings.Replace(uri, "/", "@@@", -1)
	h := md5.New()
	io.WriteString(h, repStr)
	m1 := h.Sum(nil)

	b641 := base64.StdEncoding.EncodeToString(m1)
	return b641
}
