package main

import (
	"github.com/xuliangTang/athena/athena"
	v1 "istio-app/internal/userApp/v1"
)

func main() {
	athena.Ignite().Configuration(v1.NewConfiguration()).
		Mount("", nil, v1.NewUserCtl()).
		Launch()
}
