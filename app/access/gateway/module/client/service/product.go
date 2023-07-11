package service

import (
	"context"
	"github.com/soukengo/gopkg/util/ip"
	v1 "rockimserver/apis/rockim/api/client/v1/protocol/http"
	"rockimserver/app/access/gateway/conf"
)

type ProductService struct {
	cfg *conf.Config
}

func NewProductService(cfg *conf.Config) *ProductService {
	return &ProductService{cfg: cfg}
}

func (s *ProductService) FetchConfig(ctx context.Context, in *v1.ConfigFetchRequest) (out *v1.ConfigFetchResponse, err error) {
	// todo: 先暂时这样写
	internalIp := ip.InternalIP()
	out = &v1.ConfigFetchResponse{
		Socket: &v1.Socket{
			//Tcp: &v1.Socket_TCP{Address: s.cfg.Comet.TCP.Addr},
			//Ws:  &v1.Socket_Websocket{Address: s.cfg.Comet.WebSocket.Addr},
			Tcp: &v1.Socket_TCP{Address: internalIp + ":6003"},
			Ws:  &v1.Socket_Websocket{Address: "ws://" + internalIp + ":6004/"},
		},
	}
	return
}
