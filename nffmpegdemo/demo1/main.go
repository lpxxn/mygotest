package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mygotest/nffmpegdemo/demo1/cut_video_info"
	"github.com/satori/go.uuid"
)

func main() {
	filePath, fileExist := cut_video_info.FilePathAndCheck("poldarkS05E03.mp4")
	//filePath, fileExist := cut_video_info.FilePathAndCheck("大侦探皮卡丘.mp4")
	if !fileExist {
		log.Println("file Not Exist")
		return
	}
	// strconv.FormatInt(1, 10)+
	uid, _ := uuid.NewV4()
	outName := fmt.Sprintf("%s.mp4", uid)
	beforeTime := time.Now()
	cVI := &cut_video_info.CutVideoInfo{
		InputFile:  filePath,
		OutputFile: cut_video_info.OutPutFile(outName),
		StartTime:  "00:30:20.05",
		EndTime:    "00:31:20.50",
	}
	outPut, err := cut_video_info.CutVideo(cVI)
	fmt.Println(outPut)
	fmt.Println("duration: ", time.Now().Sub(beforeTime).Seconds())
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

}
