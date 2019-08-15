package public

import (
	"api_client/api/common/configuration"
	"api_client/api/public/model"
)

type Client interface {
	Ticker(symbol configuration.Symbol) (*model.TickerRes, error)
}

type client struct {
	ticker
}

func New() Client {
	return &client{}
}
