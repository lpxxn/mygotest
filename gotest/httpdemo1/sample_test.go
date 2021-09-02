package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestNameHandler(t *testing.T) {
	student := NewStudent("li")
	handler := NameHandler(student)
	req, _ := http.NewRequest(http.MethodGet, "name", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("get name api error")
	}
	body := w.Body.String()
	if len(body) == 0 {
		t.Errorf("response is empty")
	}
	t.Logf("response body %s\n", body)
}

func TestSetName(t *testing.T) {
	student := NewStudent("li")
	handler := SetName(student)
	params := url.Values{}
	params.Set("name", "peng")

	req, err := http.NewRequest(http.MethodPost, "setName", strings.NewReader(params.Encode()))
	if err != nil {
		t.Errorf("new request error %v", err)
	}
	req.Header.Set("Content-Type", "multipart/form-data;")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("get name api error")
	}
	body := w.Body.String()
	if len(body) == 0 {
		t.Errorf("response is empty")
	}
	t.Logf("response body %s\n", body)
}

func TestValue1(t *testing.T) {
	var params = url.Values{}
	params.Add("app_id", "afadfasdfasdf")
	params.Add("method", "cccc")
	params.Add("format", "json")
	params.Add("charset", "utf-8")
	params.Add("sign_type", "sign")
	params.Add("version", "v1.0")
	params.Add("app_auth_token", "aaa-bbb-ccc")
	params.Add("extend_params", fmt.Sprintf(`{"sys_service_provider_id":"%s"}`, "uniidididdi"))
	t.Log(Encode(params))

}

func Encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func TestValue2(t *testing.T) {
	type ExtendParams struct {
		SysServiceProviderID string `json:"sys_service_provider_id"`
	}
	p := ExtendParams{SysServiceProviderID: "aaaa"}
	b, _ := json.Marshal(p)
	var params = url.Values{}
	params.Add("app_id", "afadfasdfasdf")
	params.Add("method", "cccc")
	params.Add("format", "json")
	params.Add("charset", "utf-8")
	params.Add("sign_type", "sign")
	params.Add("version", "v1.0")
	params.Add("app_auth_token", "aaa-bbb-ccc")
	params.Add("extend_params", string(b))
	t.Log(Encode(params))

	type A struct {
		Name         string
		ExtendParams *struct {
			SysServiceProviderID string `json:"sys_service_provider_id"`
		} `json:"extend_params,omitempty"`
	}
	a := &A{}
	b1, _ := json.Marshal(a)
	t.Log(string(b1))
}

func TestValue3(t *testing.T) {
	type s struct {
		ID   int64
		Name string
		Age  int
	}
	type b struct {
		*s
		ID  string
		Age string
	}
	a1 := &s{
		ID:   1,
		Name: "aa",
		Age:  1,
	}
	a2 := &b{
		s:   a1,
		ID:  "abc",
		Age: "def",
	}
	j, err := json.Marshal(a2)
	t.Log(string(j), err)
}
