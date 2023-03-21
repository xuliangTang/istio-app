package v1

import (
	"context"
	"github.com/xuliangTang/athena/athena"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
)

type IProdRepo interface {
	FindAll(context.Context, *[]*ProdModel) error
}

type ProdRepo struct{}

func NewProdRepo() *ProdRepo {
	return &ProdRepo{}
}

func (this *ProdRepo) FindAll(ctx context.Context, prods *[]*ProdModel) error {
	// 客户端构建器
	c, err := builder.NewClientBuilder("localhost:8080").WithOptions(grpc.WithInsecure()).Build()
	athena.Error(err)

	// 参数构建器
	paramBuilder := builder.NewParamBuilder()

	// api构建器
	api := builder.NewApiBuilder("getProdList", builder.ApiTypeQuery)

	return api.Invoke(ctx, c, paramBuilder, &prods)
}
