package lib

import (
	"fmt"
	"time"
)

func TimeNow() {
	fmt.Println(time.Local, "time now: ", time.Now())
}

func SetTimeZoneToChina() error {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	time.Local = loc
	return nil
}
