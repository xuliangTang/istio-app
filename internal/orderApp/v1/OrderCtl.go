package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

type OrderCtl struct {
	OrderService *OrderService `inject:"-"`
}

func NewOrderCtl() *OrderCtl {
	return &OrderCtl{}
}

func (this *OrderCtl) prods(ctx *gin.Context) any {
	var prods []*ProdModel
	athena.Error(this.OrderService.GetProds(ctx, &prods))

	return prods
}

func (this *OrderCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "prods", this.prods)
}
