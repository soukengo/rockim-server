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
	"rockimserver/app/logic/user/data/cache"
	"rockimserver/app/logic/user/data/database"
	"rockimserver/app/logic/user/server"
	"rockimserver/app/logic/user/service"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(env *conf.Env, config *discovery.Config, confServer *conf.Server, mongoConfig *mongo.Config, redisConfig *redis.Config) (*kratos.App, error) {
	client := mongo.NewClient(mongoConfig)
	userData := database.NewUserData(client)
	redisClient := redis.NewClient(redisConfig)
	cacheUserData := cache.NewUserData(redisClient)
	userRepo := data.NewUserRepo(userData, cacheUserData)
	userUseCase := biz.NewUserUseCase(userRepo)
	userService := service.NewUserService(userUseCase)
	authService := service.NewAuthService(userUseCase)
	grpcServer := server.NewGRPCServer(confServer, userService, authService)
	registrar, err := discovery.NewRegistrar(config)
	if err != nil {
		return nil, err
	}
	app := newApp(env, grpcServer, registrar)
	return app, nil
}
