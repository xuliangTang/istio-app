package v1

type OrderConfiguration struct{}

func NewOrderConfiguration() *OrderConfiguration {
	return &OrderConfiguration{}
}

func (*OrderConfiguration) InitOrderRepo() *OrderRepo {
	return NewOrderRepo()
}

func (*OrderConfiguration) InitOrderService() *OrderService {
	return NewOrderService()
}
