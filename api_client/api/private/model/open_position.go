package model

import (
	"api_client/api/common/configuration"
	"api_client/api/common/model"
	"time"

	"github.com/shopspring/decimal"
)

type OpenPositionRes struct {
	model.ResponseCommon
	Data struct {
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			Count       int `json:"count"`
		} `json:"pagination"`
		List []struct {
			PositionID   int64                `json:"positionId"`
			Symbol       configuration.Symbol `json:"symbol"`
			Side         configuration.Side   `json:"side"`
			Size         decimal.Decimal      `json:"size"`
			OrderdSize   decimal.Decimal      `json:"orderdSize"`
			Price        decimal.Decimal      `json:"price"`
			LossGain     decimal.Decimal      `json:"lossGain"`
			Leverage     decimal.Decimal      `json:"leverage"`
			LosscutPrice decimal.Decimal      `json:"losscutPrice"`
			Timestamp    time.Time            `json:"timestamp"`
		} `json:"list"`
	}
}
