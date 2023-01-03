// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package platform

import (
	"github.com/go-kratos/kratos/v2"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/conf"
	"rockim/app/logic/platform/data"
	"rockim/app/logic/platform/data/database"
	"rockim/app/logic/platform/server"
	"rockim/app/logic/platform/service"
	"rockim/pkg/component/database/mongo"
	"rockim/pkg/component/discovery"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(env *conf.Env, config *discovery.Config, confServer *conf.Server, mongoConfig *mongo.Config) (*kratos.App, error) {
	client := mongo.NewClient(mongoConfig)
	platUserData := database.NewPlatUserData(client)
	platUserRepo := data.NewPlatUserRepo(platUserData)
	platUserUseCase := biz.NewPlatUserUseCase(platUserRepo)
	platUserService := service.NewPlatUserService(platUserUseCase)
	platRoleData := database.NewPlatRoleData(client)
	platRoleRepo := data.NewPlatRoleRepo(platRoleData)
	platRoleUseCase := biz.NewPlatRoleUseCase(platRoleRepo)
	platRoleService := service.NewPlatRoleService(platRoleUseCase)
	platResourceData := database.NewPlatResourceData(client)
	platResourceRepo := data.NewPlatResourceRepo(platResourceData)
	platResourceUseCase := biz.NewPlatResourceUseCase(platResourceRepo)
	platResourceService := service.NewPlatResourceService(platResourceUseCase)
	tenantData := database.NewTenantData(client)
	tenantRepo := data.NewTenantRepo(tenantData)
	tenantUseCase := biz.NewTenantUseCase(tenantRepo)
	tenantService := service.NewTenantService(tenantUseCase)
	tenantResourceData := database.NewTenantResourceData(client)
	tenantResourceRepo := data.NewTenantResourceRepo(tenantResourceData)
	tenantResourceUseCase := biz.NewTenantResourceUseCase(tenantResourceRepo)
	tenantResourceService := service.NewTenantResourceService(tenantResourceUseCase)
	grpcServer := server.NewGRPCServer(confServer, platUserService, platRoleService, platResourceService, tenantService, tenantResourceService)
	registrar, err := discovery.NewRegistrar(config)
	if err != nil {
		return nil, err
	}
	app := newApp(env, grpcServer, registrar)
	return app, nil
}
