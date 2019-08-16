package public

type Client interface {
	Ticker
}

type client struct {
	ticker
}

func New() Client {
	return &client{}
}
