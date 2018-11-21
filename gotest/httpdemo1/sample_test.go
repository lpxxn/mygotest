package main

import (
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
