// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package group

import (
	"github.com/go-kratos/kratos/v2"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/conf"
	"rockimserver/app/logic/group/data"
	cache2 "rockimserver/app/logic/group/data/cache"
	"rockimserver/app/logic/group/data/database"
	server2 "rockimserver/app/logic/group/server"
	"rockimserver/app/logic/group/service"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
	"rockimserver/pkg/component/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(config *conf.Config, discoveryConfig *discovery.Config, serverConfig *server.Config, mongoConfig *mongo.Config, redisConfig *redis.Config, cacheConfig *cache.Config) (*kratos.App, error) {
	client := mongo.NewClient(mongoConfig)
	groupData := database.NewChatRoomData(client)
	redisClient := redis.NewClient(redisConfig)
	cacheGroupData := cache2.NewGroupData(redisClient, cacheConfig)
	groupRepo := data.NewGroupRepo(groupData, cacheGroupData)
	chatRoomMemberData := cache2.NewChatRoomMemberData(redisClient, cacheConfig)
	chatRoomMemberRepo := data.NewChatRoomMemberRepo(chatRoomMemberData)
	generator := idgen.NewMongoGenerator()
	builder := lock.NewRedisBuilder(redisClient)
	chatRoomMemberManager := biz.NewChatRoomMemberManager(groupRepo, chatRoomMemberRepo, generator)
	chatRoomUseCase := biz.NewChatRoomUseCase(groupRepo, chatRoomMemberRepo, generator, builder, chatRoomMemberManager)
	chatRoomService := service.NewChatRoomService(chatRoomUseCase)
	chatRoomMemberUseCase := biz.NewChatRoomMemberUseCase(groupRepo, chatRoomMemberRepo, builder, generator, chatRoomMemberManager)
	chatRoomMemberService := service.NewChatRoomMemberService(chatRoomMemberUseCase)
	serviceGroup := server2.NewServiceGroup(chatRoomService, chatRoomMemberService)
	grpcServer := server2.NewGRPCServer(serverConfig, serviceGroup)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(config, grpcServer, registrar)
	return app, nil
}
