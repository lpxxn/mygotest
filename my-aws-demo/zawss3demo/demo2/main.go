package main

import (
	"bytes"
	"strings"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin/json"
)

func main() {

	http.HandleFunc("/upload", ReceiveFile)

	http.ListenAndServe(":5050", nil)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	// Copy the file data to my buffer
	io.Copy(&Buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	contents := Buf.String()
	fmt.Println(contents)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	Buf.Reset()
	// do something else
	// etc write header
	return
}


func ResponseJson(w http.ResponseWriter, data interface{}) {
		rdata, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(rdata)
}