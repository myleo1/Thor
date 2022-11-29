package main

import (
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"net/http"
	"net/url"
	"strings"
)

func HandleRequest(event map[string]interface{}) (string, error) {
	targetUrl := "http://xxx"
	payload := url.Values{"to": {"leo"}, "content": {"123"}}
	req, _ := http.NewRequest("POST", targetUrl, strings.NewReader(payload.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Cookie", "")
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("http client err!"), err
	}
	return fmt.Sprintf("hello, %s!", event["content"]), nil
}

func main() {
	fc.Start(HandleRequest)
}
