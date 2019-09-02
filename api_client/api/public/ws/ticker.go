package ws

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/internal/connect"
	"api_client/api/public/ws/model"
	"encoding/json"
	"log"
)

// Ticker is receiving price data.
type Ticker interface {
	SubscribeTicker(symbol configuration.Symbol) error
	UnsubscribeTicker(symbol configuration.Symbol) error
	ReceiveTicker() <-chan *model.TickerRes
}

type ticker struct {
	conn *connect.Connection
}

func (t *ticker) SubscribeTicker(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: configuration.WebSocketCommandSubscribe,
		Channel: configuration.WebSocketChannelTicker,
		Symbol:  symbol,
	}

	return t.conn.Send(req)
}

func (t *ticker) UnsubscribeTicker(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: configuration.WebSocketCommandUnsubscribe,
		Channel: configuration.WebSocketChannelTicker,
		Symbol:  symbol,
	}

	return t.conn.Send(req)
}

func (t *ticker) ReceiveTicker() <-chan *model.TickerRes {
	c := make(chan *model.TickerRes, 10)
	go func() {
		for {
			select {
			case v := <-t.conn.Stream():
				if v == nil {
					return
				}
				log.Printf("received:%v", string(v))
				res := new(model.TickerRes)
				err := json.Unmarshal(v, res)
				if err != nil {
					// TODO:error handling
					continue
				}
				c <- res
			}
		}
	}()
	return c
}
