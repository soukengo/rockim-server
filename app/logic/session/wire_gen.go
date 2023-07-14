// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package session

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/session/biz"
	"rockimserver/app/logic/session/conf"
	"rockimserver/app/logic/session/data"
	cache2 "rockimserver/app/logic/session/data/cache"
	"rockimserver/app/logic/session/infra"
	server2 "rockimserver/app/logic/session/server"
	"rockimserver/app/logic/session/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, serverConfig *server.Config, cacheConfig *cache.Config) (*kratos.App, error) {
	manager := infra.NewCacheManager(config, logger)
	channelData := cache2.NewChannelData(manager, cacheConfig)
	channelRepo := data.NewChannelRepo(channelData)
	channelUseCase := biz.NewChannelUseCase(channelRepo)
	channelService := service.NewChannelService(channelUseCase)
	channelQueryRepo := data.NewChannelQueryRepo(channelData)
	channelQueryUseCase := biz.NewChannelQueryUseCase(channelQueryRepo)
	channelQueryService := service.NewChannelQueryService(channelQueryUseCase)
	serviceGroup := server2.NewServiceGroup(channelService, channelQueryService)
	grpcServer := server2.NewGRPCServer(serverConfig, serviceGroup)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(logger, config, grpcServer, registrar)
	return app, nil
}