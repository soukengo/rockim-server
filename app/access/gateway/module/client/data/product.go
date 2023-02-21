package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/access/gateway/module/client/biz"
)

type productRepo struct {
	ac v1.ProductAPIClient
}

func NewProductRepo(ac v1.ProductAPIClient) biz.ProductRepo {
	return &productRepo{ac: ac}
}

func (r *productRepo) FindById(ctx context.Context, productId string) (out *types.Product, err error) {
	ret, err := r.ac.Find(ctx, &v1.ProductFindRequest{Id: productId})
	if err != nil {
		return
	}
	out = ret.Product
	return
}
