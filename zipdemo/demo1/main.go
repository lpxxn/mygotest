package main

import (
	"fmt"
	"github.com/mygotest/zipdemo/utils"
)

func main() {
	zipFileName := "a.zip"
	file1, file2 := "a", "b"

	err := utils.ZipFiles(zipFileName, []string{file1, file2})
	if err != nil {
		fmt.Println(err)
		return
	}
}
