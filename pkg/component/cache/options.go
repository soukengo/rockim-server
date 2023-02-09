package cache

import "time"

const (
	defaultExpire = 0
)

type Option func(opts *Options)

type Options struct {
	expire time.Duration
	codec  Codec
}

func Default() *Options {
	return &Options{expire: defaultExpire, codec: NewJsonCodec()}
}

func (o *Options) Expire() time.Duration {
	return o.expire
}

func (o *Options) Codec() Codec {
	return o.codec
}

func (o *Options) Apply(opts ...Option) *Options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithExpire(expire time.Duration) Option {
	return func(o *Options) { o.expire = expire }
}

func WithCodec(codec Codec) Option {
	return func(o *Options) { o.codec = codec }
}
