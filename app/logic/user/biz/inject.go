package biz

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
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
