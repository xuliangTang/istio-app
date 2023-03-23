package v1

import "context"

type OrderService struct {
	OrderRepo *OrderRepo `inject:"-"`
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (this *OrderService) CreateOrder(ctx context.Context, req *OrderCreateRequest) error {
	return this.OrderRepo.Create(ctx, req)
}
