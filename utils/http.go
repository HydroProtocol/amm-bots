package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type KeyPair struct {
	Key   string
	Value string
}

var EmptyKeyPairList = []KeyPair{}
var HttpClient *http.Client

func init() {
	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 15 * time.Second,
	}

	HttpClient = &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}

func callHttp(methodType string, url string, requestBody string, params []KeyPair, headers []KeyPair) (string, error) {
	var body string
	var errorCatch error
	var buffer bytes.Buffer

	buffer.WriteString(url)
	if len(params) > 0 && !strings.HasSuffix(url, "?") {
		buffer.WriteString("?")
	}
	for i, param := range params {
		buffer.WriteString(param.Key)
		buffer.WriteString("=")
		buffer.WriteString(param.Value)
		if i < len(params)-1 {
			buffer.WriteString("&")
		}
	}

	req, err := http.NewRequest(methodType, buffer.String(), bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		errorCatch = errors.New("build request failed")
	} else {
		for _, header := range headers {
			req.Header.Set(header.Key, header.Value)
		}
		resp, err := HttpClient.Do(req)
		if err != nil {
			errorCatch = err
		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errorCatch = errors.New("read resp failed")
			}
			body = string(bodyBytes)
		}
		closeBody(resp)
	}
	return body, errorCatch
}

func Post(url string, requestBody string, params []KeyPair, headers []KeyPair) (string, error) {
	return callHttp("POST", url, requestBody, params, headers)
}

func Get(url string, requestBody string, params []KeyPair, headers []KeyPair) (string, error) {
	return callHttp("GET", url, requestBody, params, headers)
}

func Delete(url string, requestBody string, params []KeyPair, headers []KeyPair) (string, error) {
	return callHttp("DELETE", url, requestBody, params, headers)
}

func closeBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}
}

func JoinUrlPath(baseUrl string, subPath string) string {
	u, _ := url.Parse(baseUrl)
	u.Path = path.Join(u.Path, subPath)
	return u.String()
}
