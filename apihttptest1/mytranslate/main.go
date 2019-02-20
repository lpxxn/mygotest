package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	sl string
	tl string
)

func init() {
	flag.StringVar(&sl, "sl", "en", "source language")
	flag.StringVar(&tl, "tl", "zh-cn", "source language")
}

var commonHandler = map[string]string{}

func addCommonHandler(req *http.Request) {

}

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://translate.google.cn/", nil)
	if err != nil {
		panic(err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	//doc, err := html.Parse(bytes.NewReader(body))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(doc)
	var re = regexp.MustCompile(`tkk:'(\d+\.\d+)'`)
	matchTkk := re.FindAllStringSubmatch(string(body), -1)
	fmt.Println(matchTkk)
	v := string(84)

	summar := "夏日"

	summarRune := []rune(summar)
	fmt.Println(v, len(summar), summarRune)
}

//

func Io(a string) {
	aStr := []rune(a)

	Ho := "430735.1017041714"
	b := Ho

	d := []string{string(116), string(107)}
	c := "&" + strings.Join(d, "")
	d = strings.Split(b, ".")
	bNum, _ := strconv.Atoi(d[0])
	b1Num, _ := strconv.Atoi(d[1])
	aStrLen := len(aStr)
	e := make([]rune, 0)
	for  g := 0; g < aStrLen; g++ {
		k := aStr[g]
		if 128 > k {
			e = append(e, k)
		} else if 2048 > k {
			e = append(e, k>>6 | 192)
		} else {
			if 55296 == int(k&64512) && g+1 < aStrLen && 56320 == (aStr[g+1]&64512) {
				g++
				k = 65536 + ((k & 1023) << 10) + (aStr[g] & 1023)
				e = append(e, k>>18 | 240)
				e = append(e, k>>12&63 | 128)
			} else {
				e = append(e, k>>12 | 224)
				e = append(e, k>>6&63 | 128)
			}
			e = append(e,  k&63 | 128)
		}
	}
	aNum := bNum
	for _, v := range e  {
		aNum += int(v)
		aNum = Go(aNum, []rune("+-a^+6"))
	}
	aNum = Go(aNum, []rune("+-3^+b+-f"))
	// b -334485971 a -135051198  d: ["430737", "467889775"]
	aNum ^= b1Num


}

func Go(a int, b []rune) int {
	for c := 0; c < len(b) - 2; c += 3 {
		d := b[c +2]
		d = func() rune {
			if rune('a') < d {
				return b[0] - 87
			}
			return d
		}()
		d = func() rune {
			if rune('+') == b[c + 1] {
				sv := uint(a) >> uint(d)
				return rune(sv)
			}
			sv := uint(a) << uint(d)
			return rune(sv)
		}()
		a = func() int {
			if rune('+') == b[c] {
				 yd := int(d) & 4294967295
				 return a + yd
			}
			return int(a) ^ int(d)
		}()
	}
	return a
}

// fmt.Println([]rune("absdef")[2])      // Also prints 115
//fmt.Printf("%c", []rune("absdef")[2]) // Prints s

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}
