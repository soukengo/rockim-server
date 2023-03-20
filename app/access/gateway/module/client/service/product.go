package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/v1/http"
	"rockimserver/app/access/gateway/conf"
)

type ProductService struct {
	cfg *conf.Config
}

func NewProductService(cfg *conf.Config) *ProductService {
	return &ProductService{cfg: cfg}
}

func (s *ProductService) FetchConfig(ctx context.Context, in *v1.ConfigFetchRequest) (out *v1.ConfigFetchResponse, err error) {
	out = &v1.ConfigFetchResponse{
		Socket: &v1.Socket{
			Tcp: &v1.Socket_TCP{Address: s.cfg.Comet.TCP.Addr},
			Ws:  &v1.Socket_Websocket{Address: s.cfg.Comet.WebSocket.Addr},
		},
	}
	return
}
