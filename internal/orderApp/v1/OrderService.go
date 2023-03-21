package v1

import "context"

type OrderService struct {
	ProdRepo *ProdRepo `inject:"-"`
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (this *OrderService) GetProds(ctx context.Context, prods *[]*ProdModel) error {
	return this.ProdRepo.FindAll(ctx, prods)
}
