package v1

type ProdConfiguration struct{}

func NewProdConfiguration() *ProdConfiguration {
	return &ProdConfiguration{}
}

func (*ProdConfiguration) InitRepo() *ProdRepo {
	return NewProdRepo()
}

func (*ProdConfiguration) InitService() *ProdService {
	return NewProdService()
}
