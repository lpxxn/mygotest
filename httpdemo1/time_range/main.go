package time_range

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//DayTime Time of a day
type DayTime struct {
	Hour   int
	Minute int
	Second int
}

const (
	//HourInSecond one hour equals to 60*60 second
	HourInSecond = 60 * 60
	//MinuteInSecond one minute equals to 60 second
	MinuteInSecond = 60
)

//NewDayTime to construct a new instance of DayTime
func NewDayTime(h, m, s int) *DayTime {
	return &DayTime{h, m, s}
}

func (d *DayTime) String() string {
	if d.Second == 0 {
		return fmt.Sprintf("%02d:%02d", d.Hour, d.Minute)
	}
	return fmt.Sprintf("%02d:%02d:%02d", d.Hour, d.Minute, d.Second)
}

func (d *DayTime) Seconds() int {
	return HourInSecond*d.Hour + MinuteInSecond*d.Minute + d.Second
}

//Before returns true/flase according to time clock
func (d *DayTime) Before(other *DayTime) bool {
	return d.Seconds() < other.Seconds()
}

// NotBefore >=
func (d *DayTime) NotBefore(other *DayTime) bool {
	return d.Seconds() >= other.Seconds()
}

//After returns true/flase according to time clock
func (d *DayTime) After(other *DayTime) bool {
	return d.Seconds() > other.Seconds()
}

//NotAfter <=
func (d *DayTime) NotAfter(other *DayTime) bool {
	return d.Seconds() <= other.Seconds()
}

//MarshalJSON for json serialization
func (d *DayTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Seconds())
}

//UnmarshalJSON to deserialize from JSON
func (d *DayTime) UnmarshalJSON(src []byte) error {
	var seconds int
	err := json.Unmarshal(src, &seconds)
	if err != nil {
		return err
	}
	data := computeBySeonds(seconds)
	*d = *data
	return nil
}

func computeBySeonds(src int) *DayTime {
	var hour = src / HourInSecond
	var minute = (src - hour*HourInSecond) / MinuteInSecond
	var second = src - hour*HourInSecond - minute*MinuteInSecond
	return &DayTime{
		Hour:   hour,
		Minute: minute,
		Second: second,
	}
}

//Value for database
func (d *DayTime) Value() (driver.Value, error) {
	return d.String(), nil
}

//Scan for DB read
func (d *DayTime) Scan(src []byte) error {
	return d.UnmarshalJSON(src)
}

func (d *DayTime) parseBytes(src []byte) ([3]int, error) {

	//strconv.Unquote --> 去除JSON中字符串带有的引号
	unquotes, err := strconv.Unquote(string(src))
	if err != nil {
		return [3]int{0, 0, 0}, err
	}
	return parseDayTimeElements(unquotes)
}

//ParseDayTime parse from the string format
// HH:mm:ss 或者 HH 或者 HH:mm 都支持
func ParseDayTime(src string) (*DayTime, error) {
	elements, err := parseDayTimeElements(src)
	if err != nil {
		return nil, err
	}
	return &DayTime{
		Hour:   elements[0],
		Minute: elements[1],
		Second: elements[2],
	}, nil
}

func parseDayTimeElements(src string) ([3]int, error) {
	data := [3]int{0, 0, 0}
	var err error
	items := strings.Split(src, ":")
	for i := range items {
		data[i], err = strconv.Atoi(items[i])
		if err != nil {
			return data, err
		}
	}
	return data, nil
}

//DayTimeRange to express a range
type DayTimeRange struct {
	From *DayTime `json:"from"`
	To   *DayTime `json:"to"`
}

func ParseDayTimeRange(from, to string) (*DayTimeRange, error) {
	dist1, err := ParseDayTime(from)
	if err != nil {
		return nil, fmt.Errorf("parse range [from] error, param value: %s", from)
	}
	dist2, err := ParseDayTime(to)
	if err != nil {
		return nil, fmt.Errorf("parse range [to] error,param value: %s", to)
	}
	return &DayTimeRange{dist1, dist2}, nil
}

func (r *DayTimeRange) String() string {
	return fmt.Sprintf(`{from:"%s",to:"%s"}`, r.From, r.To)
}

func (r *DayTimeRange) Value() (driver.Value, error) {
	return json.Marshal(r)
}
func (r *DayTimeRange) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	return json.Unmarshal(src.([]byte), r)
}

//DayTimeRanges a slice of DayTimeRange
type DayTimeRanges []*DayTimeRange

//Value to database insert
func (rs DayTimeRanges) Value() (driver.Value, error) {
	return json.Marshal(rs)
}