package chain

import (
	"encoding/json"
	"net/http"
)

func ErrorRsp(writer http.ResponseWriter, err error) {
	ret := map[string]string{
		"message": err.Error(),
	}
	b, _ := json.Marshal(ret)

	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(400)
	writer.Write(b)
}

func SuccessRsp(writer http.ResponseWriter, result interface{}) {
	b, _ := json.Marshal(result)

	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(200)
	writer.Write(b)
}

func SuccessRspByString(writer http.ResponseWriter, ret []byte) {
	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(200)

	writer.Write(ret)
}
