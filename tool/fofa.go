package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Fofa(userApi string) string {
	url := "https://fofa.info/api/v1/search/all?&qbase64=dGl0bGU9IueZvuW6piI=&fields=link&size=10&key="

	url += userApi

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		panic(fmt.Sprintf("请求失败: %v", err))
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("内容: %s\n", string(body))
	return string(body)
}
