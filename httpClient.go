package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://ad-market.vivo.com.cn/test/sign"

	payload := strings.NewReader("requestStr=%7B%22startDate%22%3A%2220190716%22%2C%22endDate%22%3A%2220190720%22%2C%22page%22%3A1%2C%22pageSize%22%3A40%7D&apiKey=oof%2BKSv0g0k%3D")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "PostmanRuntime/7.19.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "b3e2b07e-67ba-4a8e-90cf-796c30339a63,62b8cd78-6c32-4deb-9f3f-0eced55de4ea")
	req.Header.Add("Host", "ad-market.vivo.com.cn")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "145")
	req.Header.Add("Cookie", "b_marketing_subid=5186e179172ce8b6cad6")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
