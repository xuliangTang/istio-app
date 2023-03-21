package v1

type Configuration struct{}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (*Configuration) InitRepo() *ProdRepo {
	return NewProdRepo()
}

func (*Configuration) InitService() *OrderService {
	return NewOrderService()
}
