package private

import (
	"api_client/api/common/configuration"
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"

	"github.com/shopspring/decimal"
)

type Client interface {
	Order(symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.OrderRes, error)
	CancelOrder(orderID int64) error
	OpenPositions(symbol configuration.Symbol, pageNo int) (*model.OpenPositionRes, error)
	CloseOrder(positionID int64, symbol configuration.Symbol, side configuration.Side, executionType configuration.ExecutionType, price, size decimal.Decimal) (*model.CloseOrderRes, error)
}

type client struct {
	order
	cancelOrder
	openPosition
	closeOrder
}

// New create Client instance.
func New(apiKey, secretKey string) Client {
	c := &client{}
	con := connect.New(apiKey, secretKey)
	c.order.con = con
	c.cancelOrder.con = con
	c.closeOrder.con = con
	c.openPosition.con = con

	return c
}
