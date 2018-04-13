package main

import (
	"fmt"
	"github.com/speps/go-hashids"
	"time"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
)

func main() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)


	now := time.Now()
	hd.Salt = "aa0gLG-afoDw&q=golang+url+shortener&oq=golang+url+sh&gs_l=psy-ab.3.0.0j0i22i30k1.485277.486865.0.489530.6.6.0.0.0.0.367.631.2-1j1.2.0....0...1.1.64.psy-ab..4.2.631....0.7CukeZyeQs8"
	id, _ := h.Encode([]int{int(now.Unix())})
	fmt.Println(id)

	fmt.Println(GetMD5Hash(hd.Salt))


	s := "sha1 this string"
	sh := sha1.New()
	sh.Write([]byte(s))
	sha1_hash := hex.EncodeToString(sh.Sum(nil))

	fmt.Println(s, sha1_hash)


	// Assuming 'r' is set to some inbound net/http request
	form_value := []byte("login_password")
	sha1_hash2 := fmt.Sprintf("%x", sha1.Sum(form_value))

	// Then output optionally, to test
	fmt.Println(sha1_hash2)
}


func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

