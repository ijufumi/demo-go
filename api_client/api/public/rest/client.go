package rest

// Client ...
type Client interface {
	Ticker
	Status
	OrderBooks
}

type client struct {
	ticker
	status
	orderbooks
}

// New ...
func New() Client {
	return &client{}
}
