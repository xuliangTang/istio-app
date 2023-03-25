package main

import "github.com/armon/go-socks5"

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", "0.0.0.0:9000"); err != nil {
		panic(err)
	}
}