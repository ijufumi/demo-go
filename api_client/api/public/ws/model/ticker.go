package model

import (
	"api_client/api/common/configuration"
	"github.com/shopspring/decimal"
	"time"
)

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
