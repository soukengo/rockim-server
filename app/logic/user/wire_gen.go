// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/conf"
	"rockimserver/app/logic/user/data"
	cache2 "rockimserver/app/logic/user/data/cache"
	database2 "rockimserver/app/logic/user/data/database"
	"rockimserver/app/logic/user/server"
	"rockimserver/app/logic/user/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, confServer *conf.Server, databaseConfig *database.Config, cacheConfig *cache.Config) (*kratos.App, error) {
	manager := database2.NewDatabaseManager(config)
	userData := database2.NewUserData(manager)
	cacheManager := cache2.NewCacheManager(config, logger)
	cacheUserData := cache2.NewUserData(cacheManager, cacheConfig)
	userRepo := data.NewUserRepo(userData, cacheUserData)
	generator := idgen.NewMongoGenerator()
	builder := biz.NewLockBuilder(config, logger)
	userUseCase := biz.NewUserUseCase(userRepo, generator, builder)
	userService := service.NewUserService(userUseCase)
	authCodeData := cache2.NewAuthCodeData(cacheManager, cacheConfig, logger)
	authCodeRepo := data.NewAuthCodeRepo(authCodeData)
	accessTokenData := cache2.NewAccessTokenData(cacheManager, cacheConfig, logger)
	accessTokenRepo := data.NewAccessTokenRepo(accessTokenData)
	authUseCase := biz.NewAuthUseCase(authCodeRepo, accessTokenRepo, userRepo)
	authService := service.NewAuthService(authUseCase)
	serviceGroup := server.NewServiceGroup(userService, authService)
	grpcServer := server.NewGRPCServer(confServer, serviceGroup)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(logger, config, grpcServer, registrar)
	return app, nil
}
