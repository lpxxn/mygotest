package models

import "time"

type P1Req struct {
	Name    string    `json:"name" form:"name"`
	Age     int       `json:"age" form:"age"`
	CurTime time.Time `json:"cur_time" form:"cur_time"`
}
