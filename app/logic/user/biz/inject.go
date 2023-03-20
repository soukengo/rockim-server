package biz

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(componentSet,
	NewUserUseCase,
	NewAuthUseCase,
	NewOnlineUseCase,
	NewOnlineQueryUseCase,
)

// 注册组件
var componentSet = wire.NewSet(
	lock.NewRedisBuilder,
	idgen.NewMongoGenerator)
