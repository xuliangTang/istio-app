package chain

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var midSvc string

func init() {
	midSvc = os.Getenv("MidSvc")
}

// Next 执行下一步 到底请求哪个API 是不需要知道的
func Next(headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", "http://"+midSvc, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if err != nil {
		return nil, err
	}
	rsp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	b, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode >= 400 {
		return nil, fmt.Errorf(string(b))
	}
	return b, nil
}
