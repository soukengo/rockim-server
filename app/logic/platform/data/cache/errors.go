package cache

import "rockimserver/pkg/component/cache"

// IsErrNoCache if is no cache error
func IsErrNoCache(err error) bool {
	return cache.IsErrNoCache(err)
}
