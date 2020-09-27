package main_test

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
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
