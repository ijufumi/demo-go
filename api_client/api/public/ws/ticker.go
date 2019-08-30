package ws

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/internal/connect"
	"api_client/api/public/ws/model"
)

// Ticker is receiving price data.
type Ticker interface {
	SubscribeTicker(symbol configuration.Symbol) error
	UnsubscribeTicker(symbol configuration.Symbol) error
}

type ticker struct {
	conn *connect.Connection
}

func (t *ticker) SubscribeTicker(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: "subscribe",
		Channel: "ticker",
		Symbol:  symbol,
	}

	return t.conn.Send(req)
}

func (t *ticker) UnsubscribeTicker(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: "unsubscribe",
		Channel: "ticker",
		Symbol:  symbol,
	}

	return t.conn.Send(req)
}
