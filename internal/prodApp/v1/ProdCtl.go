package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

type ProdCtl struct {
	ProdService *ProdService `inject:"-"`
}

func NewProdCtl() *ProdCtl {
	return &ProdCtl{}
}

func (this *ProdCtl) prods(ctx *gin.Context) any {
	var prods []*ProdModel
	athena.Error(this.ProdService.GetProds(ctx, &prods))

	return prods
}

func (this *ProdCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "prods", this.prods)
}
