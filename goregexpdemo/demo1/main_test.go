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
	str := `京东方B3消费机2号厅001 (325308)
京东方B3消费机2号厅002 (325309)
京东方B3消费机2号厅003 (325310)
京东方B3消费机2号厅004 (325311)
京东方B3消费机2号厅005 (325312)
京东方B3消费机2号厅006 (325313)
京东方B3消费机2号厅007 (325314)
京东方B3消费机3号厅001 (325315)
京东方B3消费机3号厅002 (325316)
京东方B3消费机3号厅003 (325317)
京东方B3消费机3号厅004 (325318)
京东方B3消费机3号厅005 (325319)
京东方B3消费机3号厅006 (325320)
京东方B3消费机3号厅007 (325321)
京东方B3消费机3号厅008 (325322)
京东方B3消费机3号厅009 (325323)
京东方B3消费机3号厅010 (325324)
京东方B3消费机3号厅011 (325325)
京东方B3消费机3号厅012 (325326)
京东方B3消费机3号厅013 (325327)
京东方B3生活区消费机001 (325487)
京东方B3生活区消费机002 (325488)
京东方B3消费机4号厅001 (325489)
京东方B5消费机001 (319083)
京东方B5消费机002 (319084)
京东方B5消费机003 (319086)
京东方B5消费机004 (319087)
京东方B5消费机005 (319088)
京东方B5消费机006 (319089)
京东方B5消费机007 (319090)
京东方B5消费机008 (319091)
京东方B5消费机009 (319092)
京东方B5消费机010 (319093)
京东方B5消费机011 (319094)
京东方B5消费机012 (319095)
京东方B5消费机013 (319096)
京东方B5消费机014 (319097)
京东方B5消费机015 (319098)
京东方B5消费机016 (319099)
京东方B5消费机017 (319100)
京东方B5消费机018 (319101)
京东方B5消费机019 (319102)
京东方B5消费机020 (319103)
京东方B5消费机021 (319104)
京东方B5消费机022 (319105)
京东方B5消费机023 (319106)
京东方B5消费机024 (319107)
京东方B5消费机025 (319108)
京东方B5消费机026 (319109)
京东方B5消费机027 (319110)
京东方B5消费机028 (319111)
京东方B5消费机029 (319113)
京东方B5消费机030 (319114)
京东方B5消费机031 (319115)
京东方B5消费机032 (319116)
京东方B5消费机033 (319117)
京东方B5消费机034 (319118)
京东方B5消费机035 (319119)
京东方B5消费机036 (319120)
京东方B5消费机037 (319121)
京东方B5消费机038 (319122)
京东方B5消费机039 (319123)
京东方B5消费机总账号 (319124)
京东方B9一号食堂001 (325683)
京东方B9一号食堂002 (325684)
京东方B9总账号 (325685)
京东方B9二号食堂001 (325686)
京东方B9二号食堂002 (325687)
京东方B9二号食堂003 (325688)
京东方B9二号食堂004 (325689)
京东方B9二号食堂005 (325690)
京东方B9二号食堂006 (325691)
京东方B9二号食堂007 (325692)
京东方B9二号食堂008 (325701)
京东方B9三号食堂001 (325702)
京东方B9三号食堂002 (325703)
京东方B9三号食堂003 (325704)
京东方B9三号食堂004 (325705)
京东方B9三号食堂005 (325706)
京东方B9三号食堂006 (325707)
京东方B9三号食堂008 (325708)
京东方B9三号食堂007 (325709)
京东方B9一号食堂003 (325710)
京东方B9生活区001 (325711)
合肥BOE医院消费机001 (318397)
合肥BOE医院消费机002 (318398)
合肥BOE医院消费机003 (318399)
合肥BOE医院消费机004 (318400)
合肥BOE医院消费机005 (318401)
合肥BOE医院消费机006 (318402)
合肥BOE医院消费机007 (318403)
合肥BOE医院消费机008 (318405)
合肥BOE医院消费机009 (318406)
合肥BOE医院消费机010 (318407)
合肥BOE医院消费机总账号 (318408)
合肥BOE医院营养餐厅 (319612)
合肥BOE医院消费机011 (321634)
合肥BOE医院消费机012 (321635)
京东方S02咖啡厅消费机001 (327800)
京东方S02消费机002 (327801)
京东方S02消费机003 (327802)
京东方S02消费机004 (327803)
京东方S02消费机005 (327804)
京东方S02消费机006 (327805)
京东方S02消费机007 (327806)
京东方S02消费机008 (327807)
京东方S02消费机009 (327808)
京东方S02生活区消费机001 (327809)
京东方S02生活区消费机002 (327810)
京东方S02消费机总账号 (327811)
京东方S02消费机010 (329291)
京东方S02消费机011 (329381)
京东方S02消费机012 (329736)
京东方S02消费机013 (334683)
京东方S02消费机014 (337471)
京东方S07消费机001 (327796)
京东方S07消费机002 (327797)
京东方S07消费机003 (327798)
京东方S07消费总账号 (327799)
京东方S07消费机004 (336134)
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

	req, err := http.NewRequest("POST", "https://bao.meiasdfn.com/patch/syncMeicanRestaurant", body)
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

	// curl --request POST 'https://baseinfo.meican.com/patch/syncMeicanRestaurant' \
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
