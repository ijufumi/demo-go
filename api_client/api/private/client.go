package private

import (
	"api_client/api/private/internal/connect"
	"os"
)

// Client ...
type Client interface {
	AccountMargin

	Orders
	Order
	ActiveOrders
	CancelOrder
	OpenPositions
	CloseOrder
}

type client struct {
	accountMargin

	orders
	order
	activeOrders
	cancelOrder
	openPositions
	closeOrder
}

// New create Client instance.
func New(apiKey, secretKey string) Client {
	c := &client{}
	con := connect.New(apiKey, secretKey)
	c.accountMargin.con = con
	c.orders.con = con
	c.order.con = con
	c.activeOrders.con = con
	c.cancelOrder.con = con
	c.closeOrder.con = con
	c.openPositions.con = con

	return c
}

// NewWithEnv ...
func NewWithEnv() Client {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("API_SECRET")

	c := &client{}
	con := connect.New(apiKey, secretKey)
	c.order.con = con
	c.activeOrders.con = con
	c.cancelOrder.con = con
	c.closeOrder.con = con
	c.openPositions.con = con

	return c
}
