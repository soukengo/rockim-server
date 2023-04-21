// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package admin

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/soukengo/gopkg/component/auth"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/access/admin/conf"
	database2 "rockimserver/app/access/admin/infra/database"
	"rockimserver/app/access/admin/infra/grpc"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/data"
	database3 "rockimserver/app/access/admin/module/manager/data/database"
	"rockimserver/app/access/admin/module/manager/service"
	biz2 "rockimserver/app/access/admin/module/tenant/biz"
	data2 "rockimserver/app/access/admin/module/tenant/data"
	database4 "rockimserver/app/access/admin/module/tenant/data/database"
	service2 "rockimserver/app/access/admin/module/tenant/service"
	server2 "rockimserver/app/access/admin/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, serverConfig *server.Config, authConfig *auth.Config, databaseConfig *database.Config) (*kratos.App, error) {
	manager := database2.NewDatabaseManager(config)
	sysUserData := database3.NewSysUserData(manager)
	sysUserRepo := data.NewSysUserRepo(sysUserData)
	authUseCase := biz.NewAuthUseCase(authConfig, sysUserRepo)
	authService := service.NewAuthService(authUseCase)
	sysRoleData := database3.NewSysRoleData(manager)
	sysRoleRepo := data.NewSysRoleRepo(sysRoleData)
	sysResourceData := database3.NewSysResourceData(manager)
	sysResourceRepo := data.NewSysResourceRepo(sysResourceData)
	sessionUseCase := biz.NewSessionUseCase(sysUserRepo, sysRoleRepo, sysResourceRepo)
	sessionService := service.NewSessionService(sessionUseCase)
	sysUserUseCase := biz.NewSysUserUseCase(sysUserRepo)
	sysUserService := service.NewSysUserService(sysUserUseCase)
	sysRoleUseCase := biz.NewSysRoleUseCase(sysRoleRepo)
	sysRoleService := service.NewSysRoleService(sysRoleUseCase)
	sysResourceUseCase := biz.NewSysResourceUseCase(sysResourceRepo)
	sysResourceService := service.NewSysResourceService(sysResourceUseCase)
	registryDiscovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		return nil, err
	}
	tenantAPIClient, err := grpc.NewTenantAPIClient(registryDiscovery)
	if err != nil {
		return nil, err
	}
	tenantRepo := data.NewTenantRepo(tenantAPIClient)
	tenantUseCase := biz.NewTenantUseCase(tenantRepo)
	tenantService := service.NewTenantService(tenantUseCase)
	sysTenantResourceData := database3.NewTenantResourceData(manager)
	sysTenantResourceRepo := data.NewSysTenantResourceRepo(sysTenantResourceData)
	sysTenantResourceUseCase := biz.NewSysTenantResourceUseCase(sysTenantResourceRepo)
	sysTenantResourceService := service.NewSysTenantResourceService(sysTenantResourceUseCase)
	managerServiceGroup := server2.NewManagerServiceGroup(authConfig, authService, sessionService, sysUserService, sysRoleService, sysResourceService, tenantService, sysTenantResourceService)
	bizTenantRepo := data2.NewTenantRepo(tenantAPIClient)
	bizAuthUseCase := biz2.NewAuthUseCase(authConfig, bizTenantRepo)
	serviceAuthService := service2.NewAuthService(bizAuthUseCase)
	databaseSysTenantResourceData := database4.NewTenantResourceData(manager)
	bizSysTenantResourceRepo := data2.NewSysTenantResourceRepo(databaseSysTenantResourceData)
	bizSessionUseCase := biz2.NewSessionUseCase(bizSysTenantResourceRepo)
	serviceSessionService := service2.NewSessionService(bizSessionUseCase)
	productAPIClient, err := grpc.NewProductAPIClient(registryDiscovery)
	if err != nil {
		return nil, err
	}
	productRepo := data2.NewProductRepo(productAPIClient)
	productUseCase := biz2.NewProductUseCase(productRepo)
	productService := service2.NewProductService(productUseCase, bizSessionUseCase)
	tenantServiceGroup := server2.NewTenantServiceGroup(authConfig, serviceAuthService, serviceSessionService, productService)
	httpServer := server2.NewHTTPServer(serverConfig, managerServiceGroup, tenantServiceGroup)
	app := newApp(logger, config, httpServer)
	return app, nil
}
