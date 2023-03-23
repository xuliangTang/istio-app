package v1

import (
	"context"
	"fmt"
	"github.com/xuliangTang/istio-dbcore-sdk/pbfiles"
	"github.com/xuliangTang/istio-dbcore-sdk/pkg/builder"
	v1 "istio-app/internal/prodApp/v1"
)

type IOrderRepo interface {
	Create(context.Context, *OrderCreateRequest) error // 创建订单
}

type OrderRepo struct {
	APIClient pbfiles.DBServiceClient `inject:"-"`
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

func (this *OrderRepo) Create(ctx context.Context, req *OrderCreateRequest) error {
	// 创建事务API
	txApi := builder.NewTxApi(ctx, this.APIClient)

	err := txApi.Tx(func(tx *builder.TxApi) error {
		psParam := builder.NewParamBuilder().Add("prodId", req.ProdId)
		ps := &v1.ProdStockModel{}
		// 开始查询库存
		err := tx.QueryForModel("getstock", psParam, ps)
		if err != nil {
			return err
		}
		// fmt.Println("当前库存为:", ps.Stock)
		if ps.Stock < req.ProdNum {
			return fmt.Errorf("库存不够了")
		}

		setStockParam := builder.NewParamBuilder().
			Add("prodId", req.ProdId).
			Add("stock", ps.Stock-req.ProdNum).
			Add("version", ps.Version)
		// 开始扣减库存
		execRet := &ExecResult{} // 增删改执行结果
		err = tx.Exec("setstock", setStockParam, execRet)
		if err != nil {
			return fmt.Errorf("扣减库存失败:%s", err.Error())
		}
		if execRet.RowsAffected == 0 {
			return fmt.Errorf("扣减库存没有成功")
		}

		// 开始创建订单
		orderCreateParam := builder.NewParamBuilder().
			Add("orderNo", req.OrderNo).
			Add("prodID", req.ProdId).
			Add("prodNum", req.ProdNum).
			Add("prodPrice", req.ProdPrice).
			Add("orderPrice", req.ProdPrice*float32(req.ProdNum)). //计算总价
			Add("userName", req.UserName)
		execRet = &ExecResult{}
		err = tx.Exec("createorders", orderCreateParam, execRet)
		if err != nil {
			return fmt.Errorf("创建主订单失败:%s", err.Error())
		}
		if execRet.RowsAffected == 0 {
			return fmt.Errorf("创建主订单没有成功")
		}
		execRet = &ExecResult{}

		// 插入子订单
		err = tx.Exec("createordersdetail", orderCreateParam, execRet)
		if err != nil {
			return fmt.Errorf("创建子订单失败:%s", err.Error())
		}
		if execRet.RowsAffected == 0 {
			return fmt.Errorf("创建子订单没有成功")
		}
		return nil
	})

	return err
}
