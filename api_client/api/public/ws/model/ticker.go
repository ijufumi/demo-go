package model

import (
	"api_client/api/common/configuration"
	"time"

	"github.com/shopspring/decimal"
)

// TickerReq is request of ticker.
type TickerReq struct {
	Command              string `json:"command"`
	Channel              string `json:"command"`
	configuration.Symbol `json:"symbol"`
}

// TickerRes is response of ticker.
type TickerRes struct {
	Channel   string               `json:"channel"`
	Ask       decimal.Decimal      `json:"ask"`
	Bid       decimal.Decimal      `json:"bid"`
	High      decimal.Decimal      `json:"high"`
	Last      decimal.Decimal      `json:"last"`
	Low       decimal.Decimal      `json:"low"`
	Symbol    configuration.Symbol `json:"symbol"`
	Timestamp time.Time            `json:"timestamp"`
	Volume    decimal.Decimal      `json:"volume"`
}
