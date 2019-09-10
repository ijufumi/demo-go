package public

// Client ...
type Client interface {
	Ticker
	Status
}

type client struct {
	ticker
	status
}

// New ...
func New() Client {
	return &client{}
}
