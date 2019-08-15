package public

import (
	"api_client/api/common/configuration"
	"api_client/api/public/internal/connect"
	"api_client/api/public/model"
	"encoding/json"
	"net/url"
)

type ticker struct {
	con connect.Connection
}

func (t ticker) Ticker(symbol configuration.Symbol) (*model.TickerRes, error) {
	param := url.Values{}

	if symbol != configuration.Symbol_NONE {
		param.Set("symbol", string(symbol))
	}

	res, err := t.con.Get(param, "/v1/ticker")
	if err != nil {
		return nil, err
	}

	tickerRes := new(model.TickerRes)
	err = json.Unmarshal(res, tickerRes)

	if err != nil {
		return nil, fmt.Errorf("[Ticker]error:%v,body:%s", err, res)
	}

	if len(tickerRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", tickerRes.Messages)
	}

	return tickerRes, nil
}
