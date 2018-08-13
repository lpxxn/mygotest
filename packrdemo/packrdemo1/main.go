package main

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

func main() {

	box := packr.NewBox("./tmpfs")

	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":3000", nil)

}
