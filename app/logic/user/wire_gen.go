// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"github.com/go-kratos/kratos/v2"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/conf"
	"rockimserver/app/logic/user/data"
	cache2 "rockimserver/app/logic/user/data/cache"
	"rockimserver/app/logic/user/data/database"
	server2 "rockimserver/app/logic/user/server"
	"rockimserver/app/logic/user/service"
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
	userData := database.NewUserData(client)
	redisClient := redis.NewClient(redisConfig)
	cacheUserData := cache2.NewUserData(redisClient, cacheConfig)
	userRepo := data.NewUserRepo(userData, cacheUserData)
	generator := idgen.NewMongoGenerator()
	builder := lock.NewRedisBuilder(redisClient)
	userUseCase := biz.NewUserUseCase(userRepo, generator, builder)
	userService := service.NewUserService(userUseCase)
	authCodeData := cache2.NewAuthCodeData(redisClient, cacheConfig)
	authCodeRepo := data.NewAuthCodeRepo(authCodeData)
	accessTokenData := cache2.NewAccessTokenData(redisClient, cacheConfig)
	accessTokenRepo := data.NewAccessTokenRepo(accessTokenData)
	authUseCase := biz.NewAuthUseCase(authCodeRepo, accessTokenRepo, userRepo)
	authService := service.NewAuthService(authUseCase)
	onlineData := cache2.NewOnlineData(redisClient, cacheConfig)
	onlineRepo := data.NewOnlineRepo(onlineData)
	onlineUseCase := biz.NewOnlineUseCase(onlineRepo, accessTokenRepo)
	onlineService := service.NewOnlineService(onlineUseCase)
	serviceGroup := server2.NewServiceGroup(userService, authService, onlineService)
	grpcServer := server2.NewGRPCServer(serverConfig, serviceGroup)
	registrar, err := discovery.NewRegistrar(discoveryConfig)
	if err != nil {
		return nil, err
	}
	app := newApp(config, grpcServer, registrar)
	return app, nil
}
