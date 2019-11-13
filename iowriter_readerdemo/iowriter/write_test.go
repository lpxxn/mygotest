package main_test

import (
	"bytes"
	"os"
	"testing"
)

func TestWrite1(t *testing.T) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile("./abc", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var writeBuf  bytes.Buffer

	for i := 0; i < 5; i++ {
		writeBuf.Reset()
		// Write some text line-by-line to file.
		writeBuf.WriteString("hello  ")
		_, err = file.Write(writeBuf.Bytes())
		if err != nil {
			t.Fatal(err)
		}
	}


	////Save file changes.
	err = file.Sync()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("File Updated Successfully.")
}
