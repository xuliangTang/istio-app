package main

import (
	"github.com/xuliangTang/athena/athena"
	orderV1 "istio-app/internal/orderApp/v1"
	prodV1 "istio-app/internal/prodApp/v1"
)

func main() {
	athena.Ignite().Configuration(prodV1.NewProdConfiguration(), orderV1.NewOrderConfiguration()).
		Mount("", nil, prodV1.NewProdCtl(), orderV1.NewOrderCtl()).
		Launch()
}
