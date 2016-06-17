package auth

import "golang.org/x/net/context"

type (
	serviceKey struct{}

	Credentials struct {
		User     string
		Password string
	}

	Server struct {
		Host   string
		Port   string
		Format string
	}

	Auth interface {
		Init(opts ...Option) error
		Options() Options
	}

	auth struct {
		opts Options
	}
)

var (
	DefaultAuth   = newAuth()
	defaultBroker = Credentials{}
	defaultApi    = Server{}
)

func newAuth(opts ...Option) Auth {
	options := Options{
		Broker: defaultBroker,
		Api:    defaultApi,
	}

	for _, o := range opts {
		o(&options)
	}

	auth := new(auth)
	auth.opts = options

	return auth
}

func (a *auth) Options() Options {
	return a.opts
}

func (a *auth) Init(opts ...Option) error {
	for _, o := range opts {
		o(&a.opts)
	}
	return nil
}

func DefaultOptions() Options {
	return DefaultAuth.Options()
}

func Init(opts ...Option) error {
	return DefaultAuth.Init(opts...)
}

func FromContext(ctx context.Context) (Auth, bool) {
	s, ok := ctx.Value(serviceKey{}).(Auth)
	return s, ok
}

func NewContext(ctx context.Context, a Auth) context.Context {
	return context.WithValue(ctx, serviceKey{}, a)
}

func New(opts ...Option) Auth {
	return newAuth(opts...)
}
