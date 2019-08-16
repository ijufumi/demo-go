package private

import (
	"api_client/api/private/internal/connect"
)

type Client interface {
	Order
	ActiveOrders
	CancelOrder
	OpenPositions
	CloseOrder
}

type client struct {
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
	c.order.con = con
	c.activeOrders.con = con
	c.cancelOrder.con = con
	c.closeOrder.con = con
	c.openPositions.con = con

	return c
}
