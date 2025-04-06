package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Fofa(url string, transport *http.Transport) string {

	var client *http.Client
	if transport != nil {
		client = &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		}
	} else {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(fmt.Sprintf("请求失败: %v", err))
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("内容: %s\n", string(body))
	return string(body)
}
