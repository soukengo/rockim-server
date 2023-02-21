package cache

import (
	xerrors "errors"
	"rockimserver/pkg/errors"
)

var (
	ErrNoCache  = errors.NotFound("CACHE_NOT_FOUND", "no cache")
	ErrNotFound = errors.NotFound("DATA_NOT_FOUND", "data not found")
)

func IsErrNoCache(err error) bool {
	return xerrors.Is(err, ErrNoCache)
}
