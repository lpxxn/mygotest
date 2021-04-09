package main_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
)

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Println("----")
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))

	//fmt.Println()
}

// /r return 回车
func TestReq(t *testing.T) {
	fmt.Println("Download Started")

	//fileUrl := "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"
	//fileUrl := "https://jsonplaceholder.typicode.com/todos/1"
	fileUrl := "https://www.10wallpaper.com/wallpaper/1366x768/1505/Stand_By_Me_Doraemon_Movie_HD_Widescreen_Wallpaper_1366x768.jpg"
	//err := DownloadFile("avatar.jpg", fileUrl)
	err := ReqAPI(fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")

	for i := 0; i < 10; i++ {
		//fmt.Printf("\r index: %d", i)
		fmt.Printf("abcdef \r def   aaabb")
		fmt.Printf("aaa\r index: %d", i)
		time.Sleep(time.Second / 2)
	}
	fmt.Println()
}

func ReqAPI(url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}

	var w int64
	var out *bytes.Buffer = bytes.NewBuffer(make([]byte, 0, 10))
	w, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	//if _, err = io.Copy(os.Stdout, io.TeeReader(resp.Body, counter)); err != nil {
	//	return err
	//}

	//out := &BWriter{}
	//w, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	//if err != nil {
	//	return err
	//}

	fmt.Print("\n")
	fmt.Println("writer: ", w)
	return nil
}

type BWriter struct {
	buff bytes.Buffer
}

func (b *BWriter) Write(p []byte) (n int, err error) {
	//b.Total += len(p)
	return b.buff.Write(p)
	//return n, nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
