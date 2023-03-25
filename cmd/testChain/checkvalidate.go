package main

import (
	"fmt"
	"istio-app/chain"
	"net/http"
)

func err() {

}
func main() {
	// 假设这一步是验证商品有效性
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		if id == "101" || id == "105" { // 代表验证通过  ..  假设就是 业务判断过程
			ret, err := chain.Next(map[string]string{ // 关键
				"id":      id, // 把ID参数带过去
				"message": fmt.Sprintf("商品ID=%s可用", id),
				"step":    "prodok",
			})
			if err != nil {
				chain.ErrorRsp(writer, err)
			} else {
				chain.SuccessRspByString(writer, ret)
			}
		} else {
			chain.ErrorRsp(writer, fmt.Errorf("该ID=\"%s\"的商品不可用", id))
		}
	})

	http.ListenAndServe(":80", nil)
}
