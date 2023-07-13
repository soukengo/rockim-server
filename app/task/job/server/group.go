package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/soukengo/gopkg/component/server/job"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/app/task/job/service"
)

type ServiceGroup struct {
	cometSrv *service.CometService
}

func NewServiceGroup(cometSrv *service.CometService) *ServiceGroup {
	return &ServiceGroup{cometSrv: cometSrv}
}

func (g *ServiceGroup) RegisterJob(srv job.Server) {
	srv.Register(v1.TaskType_COMET_DISPATCH.String(), wrapAction[v1.CometDispatchRequest](g.cometSrv.Dispatch))
}

func (g *ServiceGroup) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterCometAPIServer(srv, g.cometSrv)
}
