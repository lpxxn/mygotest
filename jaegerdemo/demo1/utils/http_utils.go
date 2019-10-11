package utils

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// --------------
var ErrResponse = errors.New("error response")

func GetRequestAndUnmarshalResponseData(endpoint string, revData interface{}, params ...func(r *http.Request)) (*http.Response, error) {
	resp, err := GetRequest(endpoint, params...)
	if err != nil {
		return resp, err
	}
	return resp, UnmarshalResponseDataToEntity(resp, revData)
}

func UnmarshalResponseDataToEntity(res *http.Response, data interface{}, checkResponse ...func(*http.Response) bool) error {
	defer res.Body.Close()
	for _, check := range checkResponse {
		if !check(res) {
			return ErrResponse
		}
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyData, data)
	if err != nil {
		return err
	}
	return nil
}

// GetRequest :
// 发送http Get请求
func GetRequest(endpoint string, params ...func(r *http.Request)) (*http.Response, error) {
	return newRequest(http.MethodGet, endpoint, nil, params...)
}

// newRequest:
// 创建http 请求
func newRequest(httpMethod string, endpoint string, data io.Reader, params ...func(r *http.Request)) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, endpoint, data)
	if err != nil {
		return nil, err
	}
	for _, p := range params {
		p(req)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

