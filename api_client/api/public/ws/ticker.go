package ws

import (
	"api_client/api/common/configuration"
)

type Ticker interface {
	Subscribe(symbol configuration.Symbol) error
	Unsubscribe(symbol configuration.Symbol) error
}

type ticker struct {
}

func (t *ticker) Subscribe(symbol configuration.Symbol) error {
	return nil
}

func (t *ticker) Unsubscribe(symbol configuration.Symbol) error {
	return nil
}
