package main

import (
	"istio-app/chain"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		msg := "恭喜，整个流程执行完毕" + request.Header.Get("message")

		chain.SuccessRsp(writer, map[string]string{
			"result": msg,
		})
	})
	http.ListenAndServe(":80", nil)
}
