package private

import (
	"api_client/api/private/internal/connect"
	"os"
)

// Client ...
type Client interface {
	AccountMargin
	AccountAssets

	Orders
	ActiveOrders
	Executions
	LastExecutions
	OpenPositions

	Order
	CancelOrder
	CloseOrder
}

type client struct {
	accountMargin
	accountAssets

	orders
	activeOrders
	executions
	lastExecutions
	openPositions

	order
	cancelOrder
	closeOrder
}

// New create Client instance.
func New(apiKey, secretKey string) Client {
	c := &client{}
	con := connect.New(apiKey, secretKey)
	c.accountMargin.con = con
	c.accountAssets.con = con

	c.orders.con = con
	c.activeOrders.con = con
	c.executions.con = con
	c.lastExecutions.con = con
	c.openPositions.con = con

	c.order.con = con
	c.cancelOrder.con = con
	c.closeOrder.con = con

	return c
}

// NewWithEnv ...
func NewWithEnv() Client {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("API_SECRET")
	return New(apiKey, secretKey)
}
