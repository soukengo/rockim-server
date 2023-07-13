package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/app/logic/platform/service"
)

type ServiceGroup struct {
	tenantSrv  *service.TenantService
	productSrv *service.ProductService
}

func NewServiceGroup(tenantSrv *service.TenantService, productSrv *service.ProductService) *ServiceGroup {
	return &ServiceGroup{tenantSrv: tenantSrv, productSrv: productSrv}
}

func (g *ServiceGroup) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterTenantAPIServer(srv, g.tenantSrv)
	v1.RegisterProductAPIServer(srv, g.productSrv)
}
