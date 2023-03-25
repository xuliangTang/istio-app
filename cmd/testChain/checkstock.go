package main

import (
	"fmt"
	"istio-app/chain"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		id := request.Header.Get("id")
		if id == "105" { // 代表验证通过  ..  假设就是 业务判断过程
			ret, err := chain.Next(map[string]string{ // 关键
				"id":      id, // 把ID参数带过去
				"message": fmt.Sprintf("商品ID=%s有库存", id),
				"step":    "stockok",
			})
			if err != nil {
				chain.ErrorRsp(writer, err)
			} else {
				chain.SuccessRspByString(writer, ret)
			}
		} else {
			chain.ErrorRsp(writer, fmt.Errorf("该ID=%s的商品没有库存了", id))
		}
	})

	http.ListenAndServe(":80", nil)
}
