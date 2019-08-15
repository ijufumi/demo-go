package model

import "api_client/api/common/model"

type CancelOrderReq struct {
	OrderID int64 `json:"orderId"`
}

type CancelOrderRes struct {
	model.ResponseCommon
}
