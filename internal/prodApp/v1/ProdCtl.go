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

func (this *ProdCtl) getProdOnSale(ctx *gin.Context) any {
	idUri := &IdUri{}
	athena.Error(ctx.BindUri(idUri))

	prod := &ProdModel{Id: int32(idUri.Id)}
	athena.Error(this.ProdService.GetProd(ctx, prod))

	return gin.H{"on_sale": prod.OnSale}
}

func (this *ProdCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "prods", this.prods)
	athena.Handle("GET", "prod/:id/onsale", this.getProdOnSale)
}
