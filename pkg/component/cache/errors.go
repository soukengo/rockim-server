package cache

import (
	xerrors "errors"
	"rockimserver/pkg/errors"
)

var (
	ErrNoCache  = xerrors.New("no cache")
	ErrNotFound = errors.NotFound("CACHE_ERROR", "data not found")
)

func IsErrNoCache(err error) bool {
	return xerrors.Is(err, ErrNoCache)
}
