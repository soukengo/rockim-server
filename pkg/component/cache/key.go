package cache

import "strings"

const (
	cacheKeySeparator = ":"
)

type Key string

func (k Key) GenKey(parts KeyParts) string {
	return strings.Join(append([]string{string(k)}, parts.parts...), cacheKeySeparator)
}

type KeyParts struct {
	parts []string
}

func Parts(parts ...string) KeyParts {
	return KeyParts{parts: parts}
}
