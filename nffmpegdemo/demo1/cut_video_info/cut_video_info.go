package cut_video_info

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	fileDir = "/Users/lipeng/Developer/hackathon/"
)

func FilePathAndCheck(name string) (string, bool) {
	fileStr := path.Join(fileDir, name)
	log.Println(fileStr)
	if _, err := os.Stat(fileStr); err != nil && !os.IsExist(err) {
		return "", false
	}
	return fileStr, true
}

func OutPutFile(name string) string {
	return path.Join(fileDir, name)
}

type CutVideoInfo struct {
	ID         int64  `db:"id" json:"id"`
	InputFile  string `db:"input_file" json:"-"`
	OutputFile string `db:"output_file" json:"-"`
	StartTime  string `db:"start_time" json:"startTime"`
	EndTime    string `db:"end_time" json:"endTime"`
	Md5Info    string `db:"md5_info" json:"md5Info"`
	Success    bool   `db:"success" json:"success"`
	Reason     string `db:"reason" json:"reason"`
}

func (CutVideoInfo) TableName() string {
	return "cut__video_info"
}

func CutVideo(cutVideoInfo *CutVideoInfo) (string, error) {
	cmdArguments := []string{"-ss",
		cutVideoInfo.StartTime,
		"-to",
		cutVideoInfo.EndTime,
		"-i",
		cutVideoInfo.InputFile,
		"-c:v",
		"libx264",
		"-c:a",
		"aac",
		"-strict",
		"experimental",
		"-b:a",
		"98k",
		cutVideoInfo.OutputFile}
	log.Println("command str: ", cmdArguments)
	cmd := exec.Command("ffmpeg", cmdArguments...)
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func testCommand(cutVideoInfo *CutVideoInfo) (string, error) {
	cmdArguments := []string{"-ss",
		cutVideoInfo.StartTime,
		"-to",
		cutVideoInfo.EndTime,
		"-i",
		cutVideoInfo.InputFile,
		"-c:v",
		"libx264",
		"-c:a",
		"aac",
		"-strict",
		"experimental",
		"-b:a",
		"98k",
		cutVideoInfo.OutputFile}
	log.Println("command str: ", cmdArguments)
	cmd := exec.Command("ffmpeg", cmdArguments...)

	var outBytes, errBytes bytes.Buffer
	cmd.Stdout = &outBytes
	cmd.Stderr = &errBytes
	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", errBytes.String())
		return "", err
	}
	return outBytes.String(), err
}
