package v1

import (
	"github.com/xuliangTang/athena/athena"
)

type OrderCtl struct {
	OrderService *OrderService `inject:"-"`
}

func NewOrderCtl() *OrderCtl {
	return &OrderCtl{}
}

func (this *OrderCtl) Build(athena *athena.Athena) {

}
