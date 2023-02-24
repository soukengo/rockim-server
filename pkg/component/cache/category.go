package cache

import "time"

const keySplit = ":"

// Category cache category
type Category struct {
	Category    string
	Expire      time.Duration
	EmptyExpire time.Duration
}

func (c *Category) Options() (opts []Option) {
	if c.Expire > 0 {
		opts = append(opts, WithExpire(c.Expire))
	}
	if c.EmptyExpire > 0 {
		opts = append(opts, WithEmptyExpire(c.EmptyExpire))
	}
	return
}

func (c *Category) GenKey(keySuffix string) string {
	return c.Category + keySplit + keySuffix
}

func (c *Category) CopyFrom(source *Category) (ret *Category) {
	ret = &Category{Category: source.Category, Expire: c.Expire, EmptyExpire: c.EmptyExpire}
	if source.Expire > 0 {
		ret.Expire = source.Expire
	}
	if source.EmptyExpire > 0 {
		ret.EmptyExpire = source.EmptyExpire
	}
	return
}
