package public

// Client ...
type Client interface {
	Ticker
}

type client struct {
	ticker
}

// New ...
func New() Client {
	return &client{}
}
