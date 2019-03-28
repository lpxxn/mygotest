package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
)

// ProxyTarget encapsulates url.URL and ProxyPath, when pass this to reverse proxy,
// the final target endpoint will be URL.String() + "/" + ProxyPath
type ProxyTarget struct {
	URL       *url.URL
	ProxyPath string
}

func main() {
	endpoint := "http://www.baidu.com/abc/order"
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(fmt.Errorf("GetProxyTarget: URL.Parse(%s), err = %v", endpoint, err))
	}
	var proxyPath string
	lastpath := u.Path
	ss := strings.Split(lastpath, "/")
	if len(ss) > 0 {
		proxyPath = ss[len(ss)-1]
		ss = ss[:len(ss)-1]
		u.Path = strings.Join(ss, "/")
	}
	pt := &ProxyTarget{}
	// recalculate the new target
	pt.URL, err = url.Parse(u.String())
	if err != nil {
		panic(fmt.Errorf("GetProxyTarget: URL.Parse(%s), err = %v", endpoint, err))
	}
	pt.ProxyPath = proxyPath
	reverseProxy := httputil.NewSingleHostReverseProxy(pt.URL)
	w := &MyResp{

	}
	reverseProxy.ServeHTTP(w, nil)
}

type MyResp struct {
	http.Response
	buffer bytes.Buffer
}

func (w *MyResp) Header() http.Header {
	return w.Response.Header
}

func (w *MyResp) Write(b []byte) (int, error) {
	return w.buffer.Write(b)
}

func (w *MyResp) WriteHeader(statusCode int) {
	w.Response.StatusCode = statusCode
}


func SingleJoiningSlash(a string, ss ...string) string {
	path.Join()
	sb := strings.Builder{}
	aslash := strings.HasSuffix(a, "/")
	sb.WriteString(a)

	for _, b := range ss {
		if len(b) == 0 {
			continue
		}
		bslash := strings.HasPrefix(b, "/")
		switch {
		case aslash && bslash:
			sb.WriteString(b[1:])
		case !aslash && !bslash:
			sb.WriteRune('/')
			sb.WriteString(b)
		default:
			sb.WriteString(b)
		}
		aslash = strings.HasSuffix(b, "/")
	}
	return sb.String()
}