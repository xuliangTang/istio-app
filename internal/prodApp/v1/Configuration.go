package v1

import (
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

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

func (*ProdConfiguration) InitApiClient() pbfiles.DBServiceClient {
	c, err := builder.NewClientBuilder("localhost:8080").
		WithOptions(grpc.WithInsecure()).Build()
	if err != nil {
		log.Fatal(err)
	}
	return c
}
