package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/soukengo/gopkg/component/transport/queue"
	"github.com/soukengo/gopkg/util/codec"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/app/task/job/biz/consts"
	"rockimserver/app/task/job/service"
)

type ServiceRegistry struct {
	cometSrv *service.CometService
}

func NewServiceRegistry(cometSrv *service.CometService) *ServiceRegistry {
	return &ServiceRegistry{cometSrv: cometSrv}
}

func (g *ServiceRegistry) RegisterQueue(srv queue.Server) {
	srv.Subscribe(consts.QueueCometDispatch, handleAndAck[v1.CometDispatchRequest](g.cometSrv.Dispatch, codec.JSON))
}

func (g *ServiceRegistry) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterCometAPIServer(srv, g.cometSrv)
}
