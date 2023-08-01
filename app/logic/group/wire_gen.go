// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package group

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/conf"
	"rockimserver/app/logic/group/data"
	cache2 "rockimserver/app/logic/group/data/cache"
	database2 "rockimserver/app/logic/group/data/database"
	"rockimserver/app/logic/group/infra"
	"rockimserver/app/logic/group/listener"
	"rockimserver/app/logic/group/server"
	"rockimserver/app/logic/group/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(logger log.Logger, config *conf.Config, discoveryConfig *discovery.Config, confServer *conf.Server, databaseConfig *database.Config, cacheConfig *cache.Config) (*kratos.App, error) {
	eventServer := server.NewEventServer(confServer, logger)
	manager := infra.NewDatabaseManager(config)
	groupData := database2.NewChatRoomData(manager)
	cacheManager := infra.NewCacheManager(config, logger)
	cacheGroupData := cache2.NewGroupData(cacheManager, cacheConfig)
	groupRepo := data.NewGroupRepo(groupData, cacheGroupData)
	groupUseCase := biz.NewGroupUseCase(groupRepo)
	groupService := service.NewGroupService(groupUseCase)
	chatRoomMemberData := cache2.NewChatRoomMemberData(cacheManager, cacheConfig)
	chatRoomMemberRepo := data.NewChatRoomMemberRepo(chatRoomMemberData)
	groupMemberData := cache2.NewGroupMemberData(cacheManager, cacheConfig)
	groupMemberRepo := data.NewGroupMemberRepo(groupMemberData)
	groupMemberUseCase := biz.NewGroupMemberUseCase(chatRoomMemberRepo, groupMemberRepo)
	groupMemberService := service.NewGroupMemberService(groupMemberUseCase)
	generator := idgen.NewMongoGenerator()
	builder := infra.NewLockBuilder(config, logger)
	chatRoomMemberManager := biz.NewChatRoomMemberManager(groupRepo, chatRoomMemberRepo, generator)
	chatRoomUseCase := biz.NewChatRoomUseCase(groupRepo, chatRoomMemberRepo, generator, builder, chatRoomMemberManager)
	chatRoomService := service.NewChatRoomService(chatRoomUseCase)
	chatRoomMemberUseCase := biz.NewChatRoomMemberUseCase(groupRepo, chatRoomMemberRepo, builder, generator, chatRoomMemberManager, eventServer)
	chatRoomMemberService := service.NewChatRoomMemberService(chatRoomMemberUseCase)
	serviceRegistry := server.NewServiceRegistry(groupService, groupMemberService, chatRoomService, chatRoomMemberService)
	grpcServer := server.NewGRPCServer(confServer, serviceRegistry)
	groupEventListener := listener.NewGroupEventListener(groupUseCase)
	listenerRegistry := server.NewListenerRegistry(groupEventListener)
	serverGroup := server.NewServerGroup(eventServer, grpcServer, serviceRegistry, listenerRegistry)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(config, serverGroup, registrar)
	return app, nil
}
