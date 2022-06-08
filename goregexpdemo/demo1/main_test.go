package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"
	"unicode/utf8"
)

func Benchmark_Reg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ok, _ := regexp.MatchString(`^ZONE\d*$`, "ZONE12122323"); ok {
		}

		//if ok, _ := regexp.MatchString(`^ZONE\d*$`, "zone12122323"); ok {
		//	//fmt.Println("zone12122323 ok")
		//}
	}
}

var re = regexp.MustCompile(`^ZONE\d*$`)

// good
func Benchmark_Reg2(b *testing.B) {
	//var re = regexp.MustCompile(`^ZONE\d*$`)

	for i := 0; i < b.N; i++ {
		if ok := re.MatchString("ZONE12122323"); ok {
			//fmt.Println("ZONE12122323 ok")
		}

		//if ok, _ := regexp.MatchString(`^ZONE\d*$`, "zone12122323"); ok {
		//	//fmt.Println("zone12122323 ok")
		//}
	}
}

func TestAlphabetAndNumber(t *testing.T) {
	re := regexp.MustCompile(`^[0-9_#=&?A-Za-z]{3,32}$`)
	str1 := "asdfWS23eade"
	if !re.MatchString(str1) {
		t.Error("not match: ", str1)
	}
	str1 = "asdf#123_WSe#ad_e"
	if !re.MatchString(str1) {
		t.Error("not match: ", str1)
	}
	str1 = "asdf#12一3_WSe#ad_e"
	if re.MatchString(str1) {
		t.Error("match error: ", str1)
	}
	str1 = "as12?32=df#123_WSe#ad_e"
	if !re.MatchString(str1) {
		t.Error("not match: ", str1)
	}
	str1 = "scene=2002?cid=1&sid=1"
	if !re.MatchString(str1) {
		t.Error("not match: ", str1)
	}
}

func TestCafeteriaName(t *testing.T) {
	var re = regexp.MustCompile(`^<.*>(.*)~.*\[(.*)_[0-9]*\]$`)
	var str = `<食堂菜单>2019-03-18~infinity[按天重复_27]`

	for i, match := range re.FindAllString(str, -1) {
		t.Logf(match, "found at index\n", i)
	}
	match := re.FindStringSubmatch(str)
	t.Log("match[1]: ", match[1])
	for i, name := range re.SubexpNames() {
		t.Logf("name : %s  value: %s \n", name, match[i])
	}
}

func TestCafeteriaName2(t *testing.T) {
	var re = regexp.MustCompile(`^<.*>(?P<menuDate>.*)~.*\[(?P<menuFlag>.*)_[0-9]*\]$`)
	var str = `<食堂菜单>2019-03-18~infinity[按天重复_27]`

	match := re.FindStringSubmatch(str)
	for i, name := range re.SubexpNames() {
		t.Logf("name : %s  value: %s \n", name, match[i])
	}
}

func TestDishName(t *testing.T) {
	var re = regexp.MustCompile(`-副本(?P<numFlag>[0-9]*)$`)
	var str = `菜品名字`

	match := re.FindStringSubmatch(str)
	if match != nil {
		for i, name := range re.SubexpNames() {
			t.Logf("name : %s  value: %s \n", name, match[i])
		}
	} else {
		t.Logf("no match")
	}
	r, s := utf8.DecodeLastRune([]byte(str))
	t.Log(r, s)
	t.Log(unicode.IsNumber(r))

	str = `菜1-副本`
	f1 := func() {
		match = re.FindStringSubmatch(str)
		t.Log(match)
		for i, name := range re.SubexpNames() {
			t.Logf("name : %s  value: %s \n", name, match[i])
		}
		if match != nil {
			t.Log(re.ReplaceAllLiteralString(str, ""))
		}
		r, s = utf8.DecodeLastRune([]byte(str))
		t.Log(r, s)
		t.Log(unicode.IsNumber(r))
	}
	f1()
	str = `菜-副本1-副本`
	f1()
	str = `菜-副本1-副本1`
	f1()

	str = `菜1212品-副本123`
	match = re.FindStringSubmatch(str)
	t.Log(match)
	if match != nil {
		t.Log(re.ReplaceAllLiteralString(str, ""))
	}
	for i, name := range re.SubexpNames() {
		t.Logf("name : %s  value: %s \n", name, match[i])
	}
	intList := []int{2, 41, 1, 5}
	sort.Ints(intList)
	t.Log(intList)

	r, s = utf8.DecodeLastRune([]byte(str))
	t.Log(r, s)
	t.Log(unicode.IsNumber(r))
	t.Log(strconv.Atoi(string(r)))
}

func TestReplaceCafeteriaName1(t *testing.T) {
	var re = regexp.MustCompile(`^<.*>(.*)~.*\[(.*)_[0-9]*\]$`)
	var str = `<食堂菜单>2019-03-18~infinity[按天重复_27]`

	if re.MatchString(str) {
		rs := re.ReplaceAllString(str, "$1 起[$2]")
		t.Log(rs)
	} else {
		t.Error("not match")
	}
}

func TestReplaceCafeteriaName2(t *testing.T) {
	var re = regexp.MustCompile(`^<.*>(.*)~.*\[(.*)_[0-9]*\]$`)
	var str = `<食堂菜单>2019-03-18~infinity[按天重复`

	if re.MatchString(str) {
		t.Error("match")
	}
}

func TestAFsdf(t *testing.T) {
	str := `/admin/restaurant/98/cafeteria/calendar/menus/sectione, /admin/restaurant/17/cafeteria/calendar/menus/sectione, /admin/restaurant/21/cafeteria/calendar/menus/sectione`
	var re = regexp.MustCompile(`/admin/restaurant/(?P<number>\d*)/cafeteria/calendar/menus`)
	m := re.FindStringSubmatch(`/admin/restaurant/98/cafeteria/calendar/menus/sectione`)
	t.Log(re.SubexpNames(), m)
	rev := map[string]struct{}{}
	match := re.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(match); i++ {
		t.Logf("value: %s \n", match[i][1])
		rev[match[i][1]] = struct{}{}
	}
	for k, _ := range rev {
		fmt.Printf("\"%s\", ", k)
	}
	arr := []string{"477", "131", "134", "165", "476", "905", "17", "98", "474", "130", "904", "21", "859", "536", "335", "132", "436", "164"}
	sort.Slice(arr, func(i, j int) bool {
		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[j])
		return a < b
	})
	fmt.Println(arr)

	fmt.Println(strings.Join(arr, ","))
}

func TestMultiple(t *testing.T) {
	str := `https:///corps/516082/settings
https:///corps/814139/settings
https:///corps/822386/settings
https:///corps/739363/settings`

	var re = regexp.MustCompile(`rps/(?P<number>\d*)/settings`)

	match := re.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(match); i++ {
		fmt.Printf("%s, ", match[i][1])
	}
	fmt.Println()
	for i := 0; i < len(match); i++ {
		fmt.Println(`{"corpNamespace": "` + match[i][1] + `","themeID": 2},`)
	}
}

func TestSetCo(t *testing.T) {
	str := `
京东方S07消费机005 (336135)`

	var re = regexp.MustCompile(` \((?P<number>\d*)\)`)
	match := re.FindAllStringSubmatch(str, -1)
	c := []int{336135, 336134, 327799, 327798, 327797, 327796, 337471, 334683, 329736, 329381, 329291, 327811, 327810, 327809, 327808, 327807, 327806, 327805, 327804, 327803, 327802, 327801, 327800, 319612, 318408, 325711, 325710, 325709, 325708, 325707, 325706, 325705, 325704, 325703, 325702, 325701, 325692, 325691, 325690, 325689, 325688, 325687, 325686, 325685, 325684, 325683, 319124, 319123, 319122, 319121, 319120, 319119, 319118, 319117, 319116, 319115, 319114, 319113, 319111, 319110, 319109, 319108, 319107, 319106, 319105, 319104, 319103, 319102, 319101, 319100, 319099, 319098, 319097, 319096, 319095, 319094, 319093, 319092, 319091, 319090, 319089, 319088, 319087, 319086, 319084, 319083, 325489, 325488, 325487, 325327, 325326, 325325, 325324, 325323, 325322, 325321, 325320, 325319, 325318, 325317, 325316, 325315, 325314, 325313, 325312, 325311, 325310, 325309, 325308}
	i := 0
	originMap := map[int]struct{}{}
	for ; i < len(match); i++ {
		fmt.Printf("%s, ", match[i][1])

		// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
		rID, err := strconv.Atoi(match[i][1])
		if err != nil {
			t.Fatal(err)
		}
		originMap[rID] = struct{}{}
	}

	fmt.Println()
	fmt.Println(len(c))
	fmt.Println(len(originMap))
	fmt.Println(len(match))
	for _, item := range c {
		if _, ok := originMap[item]; ok {
			delete(originMap, item)
		}
	}
	fmt.Println(originMap)
	for key, _ := range originMap {
		fmt.Printf("%d, ", key)

	}
	fmt.Println()
}

func TestQuote(t *testing.T) {
	str := `
米小奴炒饭                                                                                 |
 猫师妹(B7档口-食萤美食城)                                                                  |
  老坛鱼头王（B6档口-食萤美食城）`

	str = strings.ReplaceAll(str, "（", "(")
	str = strings.ReplaceAll(str, "）", ")")
	// +?  matches the previous token between one and unlimited times, as few times as possible, expanding as needed (lazy)
	var re = regexp.MustCompile(`(?m)(?P<name>.+?)(\((?P<subtitle>.*?)\)|$)`)
	match := re.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(match); i++ {
		t.Log(match[i])
	}

	re = regexp.MustCompile(`(?m)(?P<name>.+?)(\((?P<subtitle>.*?)\)|$)`)
	str = `Cream Soda (0.99)
Potato Chips (2.50)
Atlantic Salmon
猫师妹(B7档口-食萤美食城)
aaasdfas()
猫师妹(B7档口-食萤美食城)
一二三`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}

func TestQuote3(t *testing.T) {
	var re = regexp.MustCompile(`\$request.(?P<parameter>[.\w]+)`)
	str := `asdf{{$request.Name.abc}}=`
	match := re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("parameter")])

	str = `asdf$request.adf<12`
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("parameter")])

	str = `qwersd $request.asdf.asdf ==true`
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("parameter")])

	str = `qwersd $request.asdf.asdf ==true && $request.VVbb==123||$request.cdef=="adfe"`
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("parameter")])
	match2 := re.FindAllStringSubmatch(str, -1)
	idx := re.SubexpIndex("parameter")
	for _, matchItem := range match2 {
		t.Log(matchItem[idx])
	}

}
func TestQuote2(t *testing.T) {
	const replacement = ""
	var replacer = strings.NewReplacer(
		" ", replacement,
		"\r\n", replacement,
		"\r", replacement,
		"\n", replacement,
		"\v", replacement,
		"\f", replacement,
		"\u0085", replacement,
		"\u2028", replacement,
		"\u2029", replacement,
		"（", "(",
		"）", ")",
	)
	str := `
米小奴炒饭                        
`
	str = replacer.Replace(str)
	str = strings.ReplaceAll(str, "（", "(")
	str = strings.ReplaceAll(str, "）", ")")
	// +?  matches the previous token between one and unlimited times, as few times as possible, expanding as needed (lazy)
	var re = regexp.MustCompile(`(?P<name>.+?)(\((?P<subtitle>.*?)\)|$)`)
	match := re.FindStringSubmatch(str)
	nameIdx := re.SubexpIndex("name")
	subtitleIdx := re.SubexpIndex("subtitle")
	t.Log(nameIdx, subtitleIdx)
	str1 := match[re.SubexpIndex("name")]
	t.Log(str1)
	t.Log(match[re.SubexpIndex("subtitle")])
	for i := 0; i < len(match); i++ {
		t.Log(match[i])
	}

	str = ` 猫师妹(B7档口-食萤美食城)     `
	str = replacer.Replace(str)
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("name")])
	t.Log(match[re.SubexpIndex("subtitle")])

	str = `  老坛鱼头王（B6档口-食萤美食城）`
	//str = strings.TrimSpace(str)
	str = replacer.Replace(str)
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("name")])
	t.Log(match[re.SubexpIndex("subtitle")])

	str = `a`
	str = replacer.Replace(str)
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("name")])
	t.Log(match[re.SubexpIndex("subtitle")])

	//str = ""
	//str = replacer.Replace(str)
	//match = re.FindStringSubmatch(str)
	//t.Log(match[re.SubexpIndex("name")])
	//t.Log(match[re.SubexpIndex("subtitle")])

	str = `  `
	str = replacer.Replace(str)
	match = re.FindStringSubmatch(str)
	t.Log(match[re.SubexpIndex("name")])
	t.Log(match[re.SubexpIndex("subtitle")])

}

func TestSyncRestaurant(t *testing.T) {
	str := `合肥010 (318407)`
	var re = regexp.MustCompile(` \((?P<number>\d*)\)`)
	match := re.FindAllStringSubmatch(str, -1)
	i := 0
	for ; i < len(match); i++ {
		fmt.Printf("%s \n", match[i][1])

		// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
		rID, err := strconv.Atoi(match[i][1])
		if err != nil {
			t.Fatal(err)
		}
		doRequest(t, rID)
		time.Sleep(time.Second)
	}
	fmt.Println()
	fmt.Println(i)
	fmt.Println(len(match))
}
func doRequest(t *testing.T, rID int) {
	type Payload struct {
		RestaurantIDList []int `json:"restaurantIDList"`
	}

	data := Payload{
		RestaurantIDList: []int{rID},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://bao.meiasdfn.com/patch/syncabcRestaurant", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Clientid", "")
	req.Header.Set("Authorization", "bearer ")
	req.Header.Set("Clientsecret", "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		// handle err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(respBody))
}

// https://mholt.github.io/curl-to-go/

func TestMultiple2(t *testing.T) {
	str := `
r/edit/333740
r/edit/333723
r/edit/333741
`

	var re = regexp.MustCompile(`dit/(?P<number>\d*)`)

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl --request POST 'https://com/patch/syncabcRestaurant' \
	// --header 'clientID: TMorVol3uXnalyM7J9s5MMHZdn8HgoM' \
	// --header 'Authorization: bearer 6dWAtHFQkCitNFNjlrKKzhCBNo2KKI2' \
	// --header 'clientSecret: hQaauYWVcZsJR4zEXMdFY4ogo7lsQOT' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	// 	"restaurantIDList":[325320, 325321, 325322, 325323, 325324]
	// }'
	//

	match := re.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(match); i++ {
		fmt.Printf("%s, ", match[i][1])
	}
	fmt.Println()
	fmt.Println(len(match))
	for i := 0; i < len(match); i++ {
		fmt.Println(`{"corpNamespace": "` + match[i][1] + `","themeID": 2},`)
	}
}

func TestMultiple3(t *testing.T) {
	str := `

1.餐厅名称：东护岸-霞浦周一预订
  餐厅链接：https://om/r/edit/347572

2.餐厅名称：东护岸-霞浦周二
  餐厅链接：https://om/r/edit/347573

3.餐厅名称：东护岸-霞浦周三
  餐厅链接：https://om/r/edit/347574

4.餐厅名称：东护岸-霞浦周四
  餐厅链接：https://om/r/edit/347575

5.餐厅名称：东护岸-霞浦周五
  餐厅链接：https://om/r/edit/347577

6.餐厅名称：东护岸-霞浦周六
  餐厅链接：https://om/r/edit/347578

7.餐厅名称：东护岸-霞浦周日
  餐厅链接：https://om/r/edit/347579

8.餐厅名称：办公室-霞浦周一
  餐厅链接：https://om/r/edit/347580

9.餐厅名称：办公室-霞浦周二
  餐厅链接：https://om/r/edit/347581

10.餐厅名称：办公室-霞浦周
   餐厅链接：https:/com/r/edit/347582

11.餐厅名称：办公室-霞浦周
   餐厅链接：https:/com/r/edit/347583

12.餐厅名称：办公室-霞浦周
   餐厅链接：https:/com/r/edit/347584

13.餐厅名称：办公室-霞浦周
   餐厅链接：https:/com/r/edit/347585

14.餐厅名称：办公室-霞浦周
   餐厅链接：https:/com/r/edit/347586


`

	var re = regexp.MustCompile(`dit/(?P<number>\d*)`)

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl --request POST 'https:///patch/syncabcRestaurant' \
	// --header 'clientID: TMorVol3uXnalyM7J9s5MMHZdn8HgoM' \
	// --header 'Authorization: bearer 6dWAtHFQkCitNFNjlrKKzhCBNo2KKI2' \
	// --header 'clientSecret: hQaauYWVcZsJR4zEXMdFY4ogo7lsQOT' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	// 	"restaurantIDList":[325320, 325321, 325322, 325323, 325324]
	// }'
	//

	match := re.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(match); i++ {
		fmt.Printf("%s, ", match[i][1])
	}
	fmt.Println()
	fmt.Println(len(match))
	for i := 0; i < len(match); i++ {
		fmt.Println(`{"corpNamespace": "` + match[i][1] + `","themeID": 2},`)
	}
}

type Base struct {
	ID string
	W  int
	H  int
}

type Banner struct {
	Base
	Type int
}

func (b *Banner) CheckAttr(s *SImp) bool {
	return b.Type == s.Ban.A
}

type Video struct {
	Base
	Length int
}

func (v *Video) CheckAttr(s *SImp) bool {
	return v.Length > s.Vi.B && v.H < s.Tp
}

type SImp struct {
	Ban *struct {
		A int
	}
	Vi *struct {
		B int
	}

	Tp int
}

type verify interface {
	CheckAttr(*SImp) bool
}

type Creatives struct {
	Banner map[string][]verify
	Video  map[string][]verify
}

func TestBV(t *testing.T) {
	idx1 := "v"
	info := &Creatives{
		Banner: map[string][]verify{idx1: []verify{&Banner{}, &Banner{}}},
		Video:  map[string][]verify{idx1: []verify{&Video{}, &Video{}}},
	}
	imp := &SImp{
		Ban: &struct{ A int }{A: 1},
		Vi:  &struct{ B int }{B: 2},
	}
	creatives := []verify{}
	if imp.Ban != nil {
		creatives = info.Banner[idx1]
	} else if imp.Vi != nil {
		creatives = info.Video[idx1]
	}
	for _, item := range creatives {
		if item.CheckAttr(imp) {
			t.Log("aaaaaa")
		} else {
			t.Log("bbbb")
		}
	}
}

type PriceWrapper struct {
	ID    int64
	Price int64
}

func TestA(t *testing.T) {
	a := []*PriceWrapper{&PriceWrapper{
		ID:    1,
		Price: 2,
	}, &PriceWrapper{
		ID:    3,
		Price: 4,
	}}
	for i := range a {
		fmt.Println(a[i].ID)
	}
}
