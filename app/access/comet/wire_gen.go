// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package comet

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/infra/grpc"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/data"
	"rockimserver/app/access/comet/module/client/service"
	biz2 "rockimserver/app/access/comet/module/server/biz"
	data2 "rockimserver/app/access/comet/module/server/data"
	socket2 "rockimserver/app/access/comet/module/server/data/socket"
	service2 "rockimserver/app/access/comet/module/server/service"
	grpc2 "rockimserver/app/access/comet/server/grpc"
	"rockimserver/app/access/comet/server/socket"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, serverConfig *server.Config, protocol *conf.Protocol) (*kratos.App, error) {
	registryDiscovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		return nil, err
	}
	onlineAPIClient, err := grpc.NewOnlineAPIClient(registryDiscovery)
	if err != nil {
		return nil, err
	}
	onlineRepo := data.NewOnlineRepo(onlineAPIClient)
	channelUseCase := biz.NewChannelUseCase(config, onlineRepo)
	channelService := service.NewClientService(channelUseCase)
	socketServer := socket.NewSocketServer(serverConfig, protocol, channelService)
	channelManager := socket2.NewChannelManager(logger, socketServer)
	channelRepo := data2.NewChannelRepo(channelManager)
	bizChannelUseCase := biz2.NewChannelUseCase(channelRepo)
	serviceChannelService := service2.NewChannelService(bizChannelUseCase)
	serviceGroup := grpc2.NewServiceGroup(serviceChannelService)
	grpcServer := grpc2.NewGRPCServer(serverConfig, serviceGroup)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(logger, config, socketServer, grpcServer, registrar)
	return app, nil
}
