package ticker

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/internal/connect"
	"api_client/api/public/ws/ticker/model"
	"encoding/json"
	"log"
)

type Client interface {
	Subscribe(symbol configuration.Symbol) error
	Unsubscribe(symbol configuration.Symbol) error
	Receive() <-chan *model.TickerRes
}

type client struct {
	conn *connect.Connection
}

func New() Client {
	c := &client{
		conn: connect.New(),
	}
	return c
}

func (c *client) Subscribe(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: configuration.WebSocketCommandSubscribe,
		Channel: configuration.WebSocketChannelTicker,
		Symbol:  symbol,
	}

	return c.conn.Send(req)
}

func (c *client) Unsubscribe(symbol configuration.Symbol) error {
	req := model.TickerReq{
		Command: configuration.WebSocketCommandUnsubscribe,
		Channel: configuration.WebSocketChannelTicker,
		Symbol:  symbol,
	}

	return c.conn.Send(req)
}

func (c *client) Receive() <-chan *model.TickerRes {
	stream := make(chan *model.TickerRes, 10)
	go func() {
		for {
			select {
			case v := <-c.conn.Stream():
				if v == nil {
					return
				}
				log.Printf("received:%v", string(v))
				res := new(model.TickerRes)
				err := json.Unmarshal(v, res)
				if err != nil {
					log.Printf("[Ticker] unmarshal error:%v", err)
					continue
				}
				stream <- res
			}
		}
	}()
	return stream
}
