package cache

import (
	"github.com/soukengo/gopkg/component/cache"
)

const (
	keyUser        cache.Key = "user"
	keyUserAccount cache.Key = "user.account"

	keyAuthCode         cache.Key = "auth.code"
	keyAuthCodeReverse  cache.Key = "auth.code.reverse"
	keyAuthToken        cache.Key = "auth.token"
	keyAuthTokenReverse cache.Key = "auth.token.reverse"

	keyOnline cache.Key = "online"
)
