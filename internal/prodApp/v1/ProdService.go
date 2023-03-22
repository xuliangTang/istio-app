package v1

import "context"

type ProdService struct {
	ProdRepo *ProdRepo `inject:"-"`
}

func NewProdService() *ProdService {
	return &ProdService{}
}

func (this *ProdService) GetProds(ctx context.Context, prods *[]*ProdModel) error {
	return this.ProdRepo.FindAll(ctx, prods)
}
