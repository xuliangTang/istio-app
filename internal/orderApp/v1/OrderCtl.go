package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"net/http"
)

type OrderCtl struct {
	OrderService *OrderService `inject:"-"`
}

func NewOrderCtl() *OrderCtl {
	return &OrderCtl{}
}

func (this *OrderCtl) create(ctx *gin.Context) (v athena.Void) {
	req := &OrderCreateRequest{}
	athena.Error(ctx.BindJSON(req))
	athena.Error(this.OrderService.CreateOrder(ctx, req))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return
}

func (this *OrderCtl) Build(athena *athena.Athena) {
	athena.Handle(http.MethodPost, "/order", this.create)
}
