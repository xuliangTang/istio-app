package main

import (
	"github.com/xuliangTang/athena/athena"
	v1 "istio-app/internal/prodApp/v1"
)

func main() {
	athena.Ignite().Configuration(v1.NewProdConfiguration()).
		Mount("", nil, v1.NewProdCtl()).
		Launch()
}
