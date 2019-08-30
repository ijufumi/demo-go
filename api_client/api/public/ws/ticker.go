package ws

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/internal/connect"
	"api_client/api/public/ws/model"
)

// Ticker is receiving price data.
type Ticker interface {
	Subscribe(symbol configuration.Symbol) error
	Unsubscribe(symbol configuration.Symbol) error
}

type ticker struct {
	conn *connect.Connection
}

func (t *ticker) Subscribe(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: "subscribe",
		Channel: "ticker",
		Symbol:  symbol,
	}

	return t.conn.Send(req)
}

func (t *ticker) Unsubscribe(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: "unsubscribe",
		Channel: "ticker",
		Symbol:  symbol,
	}

	return t.conn.Send(req)
}
