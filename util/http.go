package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 读取 HTTP 响应体
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 输出 HTTP 响应体
	fmt.Println(string(body))
	return string(body)
}
