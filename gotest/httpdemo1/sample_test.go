package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNameHandler(t *testing.T) {
	student := NewStudent("li")
	handler := NameHandler(student)
	req, _ := http.NewRequest("GET", "name", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("get name api error")
	}
	body := w.Body.String()
	if len(body) == 0 {
		t.Errorf("response is empty")
	}
	t.Logf("response body %s", body)
}
