package cache

import "errors"

var (
	ErrNoCache = errors.New("no cache")
)

func IsErrNoCache(err error) bool {
	return errors.Is(err, ErrNoCache)
}
