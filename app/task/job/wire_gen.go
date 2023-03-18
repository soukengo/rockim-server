// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package job

import (
	"github.com/go-kratos/kratos/v2"
	"rockimserver/app/task/job/biz"
	"rockimserver/app/task/job/conf"
	"rockimserver/app/task/job/data"
	"rockimserver/app/task/job/data/grpc"
	server2 "rockimserver/app/task/job/server"
	"rockimserver/app/task/job/task"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/component/mq"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, serverConfig *server.Config, mqConfig *mq.Config) (*kratos.App, error) {
	registryDiscovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		return nil, err
	}
	pushManager, err := grpc.NewPushManager(config, registryDiscovery)
	if err != nil {
		return nil, err
	}
	pushRepo := data.NewPushRepo(pushManager)
	messagePushUseCase := biz.NewMessagePushUseCase(pushRepo)
	messageTask := task.NewMessageTask(messagePushUseCase)
	serviceGroup := server2.NewServiceGroup(messageTask)
	jobServer, err := server2.NewJobServer(serverConfig, mqConfig, serviceGroup, logger)
	if err != nil {
		return nil, err
	}
	app := newApp(logger, config, jobServer)
	return app, nil
}
