package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:4444/oauth2/introspect"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("authorization", "Bearer 4LI5VX_m_i7gg8B9MoPfHz9D-0qIBjs8Pc_6vsGaSHg.vIautJj5GIuBIQHwlD97mzGK-mcpkttWbicm3vsWI8E")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
