package v1

import (
	"context"
	"errors"
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
)

type IProdRepo interface {
	FindAll(context.Context, *[]*ProdModel) error
	FindOne(context.Context, *ProdModel) error
}

type ProdRepo struct {
	APIClient pbfiles.DBServiceClient `inject:"-"`
}

func NewProdRepo() *ProdRepo {
	return &ProdRepo{}
}

func (this *ProdRepo) FindAll(ctx context.Context, prods *[]*ProdModel) error {
	// 参数构建器
	paramBuilder := builder.NewParamBuilder()

	// api构建器
	api := builder.NewApiBuilder("getProdList", builder.ApiTypeQuery)

	return api.Invoke(ctx, this.APIClient, paramBuilder, &prods)
}

func (this *ProdRepo) FindOne(ctx context.Context, prod *ProdModel) error {
	paramBuilder := builder.NewParamBuilder().Add("prodId", prod.Id)
	api := builder.NewApiBuilder("getProdDetail", builder.ApiTypeQuery)

	prodList := make([]*ProdModel, 0)
	err := api.Invoke(ctx, this.APIClient, paramBuilder, &prodList)
	if err != nil {
		return err
	}

	if len(prodList) == 0 {
		return errors.New("no such prod")
	}

	// prod = prodList[0]
	*prod = *prodList[0]

	return nil
}
