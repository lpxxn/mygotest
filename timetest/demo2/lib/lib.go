package lib

import (
	"fmt"
	"time"
)

func TimeNow() {
	fmt.Println(time.Local, "time now: ", time.Now())
}
