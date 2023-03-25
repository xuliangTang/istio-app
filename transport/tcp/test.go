package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/proxy"
	"net"
	"os"
	"time"
)

func main() {
	dialer, err := proxy.SOCKS5("tcp", "110.41.142.160:31290", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "proxy connection error:", err)
		os.Exit(1)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	})

	ctx := context.Background()
	rdb.Set(ctx, "name", "txl", time.Second*10)
	fmt.Println(rdb.Get(ctx, "name"))
}
