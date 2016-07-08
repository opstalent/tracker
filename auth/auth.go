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

func FromContext(ctx context.Context) (Auth, bool) {
	s, ok := ctx.Value(serviceKey{}).(Auth)
	return s, ok
}

func New(ctx context.Context, username, password, host, port, format string) context.Context {
	return context.WithValue(ctx, serviceKey{}, authorize(username, password, host, port, format))
}

func authorize(username, password, host, port, format string) Auth {
	api := Server{
		host,
		port,
		format,
	}
	broker := Credentials{
		username,
		password,
	}

	return newAuth(Api(api), Broker(broker))
}
