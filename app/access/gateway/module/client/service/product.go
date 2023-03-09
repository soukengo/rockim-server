package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/http/v1/product"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) FetchConfig(ctx context.Context, in *v1.ConfigFetchRequest) (out *v1.ConfigFetchResponse, err error) {
	out = &v1.ConfigFetchResponse{
		Socket: &v1.Socket{
			Tcp: &v1.Socket_TCP{Address: "10.22.0.48:6003"},
			Ws:  &v1.Socket_Websocket{Address: "ws://10.22.0.48:6004/"},
		},
	}
	return
}
