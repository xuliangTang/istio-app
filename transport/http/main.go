package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRsp(request *http.Request) (*http.Response, error) {
	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *request
	// 构建roundTrip
	rsp, err := transport.RoundTrip(outReq)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func copyHeader(writer http.ResponseWriter, header map[string][]string) {
	for key, value := range header {
		for _, v := range value {
			writer.Header().Add(key, v)
		}
	}
}
func copyRspBody(writer http.ResponseWriter, rsp *http.Response) {
	// 写入http statusCode 状态码
	writer.WriteHeader(rsp.StatusCode)
	io.Copy(writer, rsp.Body) // 写入body到writer中
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(fmt.Sprintf("收到请求: method:%s "+
			"Host:%s Path:%s", request.Method, request.Host, request.URL.Path))
		// 第一步 获取响应
		rsp, err := getRsp(request)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadGateway)
			return
		}
		defer rsp.Body.Close()
		// 复制一份头
		copyHeader(writer, rsp.Header)
		// 拷贝响应内容
		copyRspBody(writer, rsp)
	})

	http.ListenAndServe(":8080", nil)
}
