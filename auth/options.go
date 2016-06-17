package auth

import "golang.org/x/net/context"

type (
	Option  func(*Options)
	Options struct {
		Api     Server
		Broker  Credentials
		Context context.Context
	}
)

func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Api(b Server) Option {
	return func(o *Options) {
		o.Api = b
	}
}

func Broker(b Credentials) Option {
	return func(o *Options) {
		o.Broker = b
	}
}
