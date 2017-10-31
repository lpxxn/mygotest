package thrlib

import "fmt"

type ThadData struct {
	Host string `json:"host"`
}

func ConnectSql() {
	fmt.Println("connect success")
}
