package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuthUseCase, NewSessionUseCase, NewProductUseCase)
