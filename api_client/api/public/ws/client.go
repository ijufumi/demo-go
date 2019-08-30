package ws

import "api_client/api/public/ws/internal/connect"

type Client interface {
	Ticker
}

type client struct {
	ticker
}

func New() Client {
	c := &client{}
	conn := connect.New()
	c.ticker.conn = conn
	return c
}
