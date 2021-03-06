package main

import (
	"flag"
	"fmt"
	"github.com/robertkrimen/otto"
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

	btv := -805041152 & 0xffffffff
	fmt.Println(btv)
	fmt.Println(btv)


	fmt.Println("go func rev: ", Go(477364483, []rune("+-a^+6")))
	vm := otto.New()
	ioFunc := `
	GO2 = function(a, b) {
		for (var c = 0; c < b.length - 2; c += 3) {
			var d = b.charAt(c + 2);
			d = "a" <= d ? d.charCodeAt(0) - 87 : Number(d);
			d = "+" == b.charAt(c + 1) ? a >>> d : a << d;
			a = "+" == b.charAt(c) ? a + d & 4294967295 : a ^ d
		}
		return a
	}
	ioRev = GO2(477364483, "+-a^+6");
	console.log('------******-----:\n', ioRev);
	console.log(-805041152 & 4294967295);
	console.log('end');
	test = 'abc';
	`
	// 接收和是最后一个变量值
	ioRv, err := vm.Run(ioFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println("begin js func test")

	fmt.Println(ioRv, "\n", -805041152 & 4294967295)
	if r2, err := vm.Run("revgo2 = GO2;"); err != nil {
		panic(err)
	} else {
		fmt.Println("revgo2: ", r2)
	}


	IOFunc := `
	var Fo = function(a) {
		return function() {
			return a
		}
	}, Go = function(a, b) {
		for (var c = 0; c < b.length - 2; c += 3) {
			var d = b.charAt(c + 2);
			d = "a" <= d ? d.charCodeAt(0) - 87 : Number(d);
			d = "+" == b.charAt(c + 1) ? a >>> d : a << d;
			a = "+" == b.charAt(c) ? a + d & 4294967295 : a ^ d
		}
		return a
	}, Io = function(a, b) {
		//if (null !== Ho)
		//	var b = Ho;
		//else {
		//	b = Fo(String.fromCharCode(84));
		//	var c = Fo(String.fromCharCode(75));
		//	b = [b(), b()];
		//	b[1] = c();
		//	b = (Ho = window[b.join(c())] || "") || ""
		//}
		var d = Fo(String.fromCharCode(116));
		c = Fo(String.fromCharCode(107));
		d = [d(), d()];
		d[1] = c();
		c = "&" + d.join("") + "=";
		d = b.split(".");
		b = Number(d[0]) || 0;
		for (var e = [], f = 0, g = 0; g < a.length; g++) {
			var k = a.charCodeAt(g);
			128 > k ? e[f++] = k : (2048 > k ? e[f++] = k >> 6 | 192 : (55296 == (k & 64512) && g + 1 < a.length && 56320 == (a.charCodeAt(g + 1) & 64512) ? (k = 65536 + ((k & 1023) << 10) + (a.charCodeAt(++g) & 1023),
			e[f++] = k >> 18 | 240,
			e[f++] = k >> 12 & 63 | 128) : e[f++] = k >> 12 | 224,
			e[f++] = k >> 6 & 63 | 128),
			e[f++] = k & 63 | 128)
		}
		a = b;
		for (f = 0; f < e.length; f++)
			a += e[f],
			a = Go(a, "+-a^+6");
		a = Go(a, "+-3^+b+-f");
		a ^= Number(d[1]) || 0;
		0 > a && (a = (a & 2147483647) + 2147483648);
		a %= 1E6;
		return c + (a.toString() + "." + (a ^ b))
	};
	ioRev = Io("summe", "431091.2043743068");
	console.log(ioRev)
	`
	if r2, err := vm.Run(IOFunc); err != nil {
		panic(err)
	} else {
		fmt.Println("revgo2: ", r2)
	}

	summar := "夏日"

	summarRune := []rune(summar)
	fmt.Println(v, len(summar), summarRune)
	Io(summar)
}

//

func Io(a string) string {
	aStr := []rune(a)

	Ho := "430755.938294120"
	b := Ho

	d := []string{string(116), string(107)}
	c := "&" + strings.Join(d, "") + "="
	d = strings.Split(b, ".")
	bNum, _ := strconv.Atoi(d[0])
	b1Num, _ := strconv.Atoi(d[1])
	aStrLen := len(aStr)
	e := make([]rune, 0)
	for g := 0; g < aStrLen; g++ {
		k := aStr[g]
		if 128 > k {
			e = append(e, k)
		} else if 2048 > k {
			e = append(e, k>>6|192)
		} else {
			if 55296 == int(k&64512) && g+1 < aStrLen && 56320 == (aStr[g+1]&64512) {
				g++
				k = 65536 + ((k & 1023) << 10) + (aStr[g] & 1023)
				e = append(e, k>>18|240)
				e = append(e, k>>12&63|128)
			} else {
				e = append(e, k>>12|224)
				e = append(e, k>>6&63|128)
			}
			e = append(e, k&63|128)
		}
	}
	aNum := bNum
	for _, v := range e {
		aNum += int(v)
		aNum = Go(aNum, []rune("+-a^+6"))
	}
	aNum = Go(aNum, []rune("+-3^+b+-f"))
	aNum ^= b1Num
	if 0 > aNum {
		aNum = (aNum & 2147483647) + 2147483648
	}
	aNum %= 1E6
	return fmt.Sprintf("%s%d.%d", c, aNum, (aNum ^ bNum))

}

func Go(a int, b []rune) int {
	for c := 0; c < len(b)-2; c += 3 {
		d := b[c+2]
		d = func() rune {
			if rune('a') <= d {
				return d - 87
			}
			n, _ := strconv.Atoi(string(d))
			return rune(n)
		}()
		d = func() rune {
			if rune('+') == b[c+1] {
				sv := uint(a) >> uint(d)
				return rune(sv)
			}
			sv := uint(a) << uint(d)
			return rune(sv)
		}()
		a = func() int {
			if rune('+') == b[c] {
				yd := int(d) & 4294967295// - 4294967296
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
