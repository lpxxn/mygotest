package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var re = regexp.MustCompile(`(\+\d{2,3}1\d{10})|(1\d{10})|(\d{3,4}\-\d{7,8})|(\d{7,8})`)
	var str = `电9875 话：122323894851asdf
6543165
6216291
15117950565
我的电话+8615117958684好了
asdfasdf：010-986435145
个-大哥21：哥1223223
asdf0312-62458754asdf
asdfcv1223asdfsadf
1221323
223-21323asdf123asdf1212
asdf一一  +86151145879962`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println("match: ", match, "found at index", i)
	}
	fmt.Println("--------")
	strs := strings.Split(str, "\n")
	for _, v := range strs {
		fs := re.FindString(v)
		fmt.Println("find str: ", fs)
	}
}
