package main_test

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
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

func TestMultiple2(t *testing.T) {
	str := `
r/edit/333740
r/edit/333723
r/edit/333741
`

	var re = regexp.MustCompile(`dit/(?P<number>\d*)`)

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
