package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func request(reqUrl string) {
	proxy, _ := url.Parse("http://110.41.142.160:31190")

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: time.Second * 3,
	}
	rsp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	b, _ := io.ReadAll(rsp.Body)
	fmt.Println(string(b))

}
func main() {
	request("http://aliyun-test/")
}
