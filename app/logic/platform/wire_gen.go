// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package platform

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/conf"
	"rockimserver/app/logic/platform/data"
	cache2 "rockimserver/app/logic/platform/data/cache"
	database2 "rockimserver/app/logic/platform/data/database"
	"rockimserver/app/logic/platform/server"
	"rockimserver/app/logic/platform/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, confServer *conf.Server, databaseConfig *database.Config, cacheConfig *cache.Config) (*kratos.App, error) {
	manager := database2.NewDatabaseManager(config)
	tenantData := database2.NewTenantData(manager)
	tenantRepo := data.NewTenantRepo(tenantData)
	tenantUseCase := biz.NewTenantUseCase(tenantRepo)
	tenantService := service.NewTenantService(tenantUseCase)
	productData := database2.NewProductData(manager)
	cacheManager := cache2.NewCacheManager(config, logger)
	cacheProductData := cache2.NewProductData(cacheManager, cacheConfig)
	productRepo := data.NewProductRepo(productData, cacheProductData)
	productUseCase := biz.NewProductUseCase(productRepo)
	productService := service.NewProductService(productUseCase)
	serviceGroup := server.NewServiceGroup(tenantService, productService)
	grpcServer := server.NewGRPCServer(confServer, serviceGroup)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(logger, config, grpcServer, registrar)
	return app, nil
}
