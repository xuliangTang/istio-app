package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Success"))
	})
	http.ListenAndServe(":80", nil)
}
