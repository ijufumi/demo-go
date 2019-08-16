package configuration

var Debug = true

type ExecutionType string

const (
	ExecutionType_MARKET = ExecutionType("MARKET")
	ExecutionType_LIMIT  = ExecutionType("LIMIT")
)

type Side string

func (s Side) Opposite() Side {
	if s == Side_BUY {
		return Side_SELL
	}
	return Side_BUY
}

const (
	Side_BUY  = Side("BUY")
	Side_SELL = Side("SELL")
)

type Symbol string

const (
	Symbol_BTC    = Symbol("BTC")
	Symbol_ETH    = Symbol("ETH")
	Symbol_BCH    = Symbol("BCH")
	Symbol_LTC    = Symbol("LTC")
	Symbol_XRP    = Symbol("XRP")
	Symbol_BTCJPY = Symbol("BTC_JPY")
	Symbol_ETHJPY = Symbol("ETH_JPY")
	Symbol_BCHJPY = Symbol("BCH_JPY")
	Symbol_LTCJPY = Symbol("LTC_JPY")
	Symbol_XRPJPY = Symbol("XRP_JPY")
	Symbol_NONE   = Symbol("")
)

type OrderType string

const (
	OrderType_NORMAL  = OrderType("NORMAL")
	OrderType_LOSSCUT = OrderType("LOSSCUT")
)

type SettleType string

const (
	SettleType_OPEN  = SettleType("OPEN")
	SettleType_CLOSE = SettleType("CLOSE")
)

type OrderStatus string

const (
	OrderStatus_WAITING    = OrderStatus("WAITING")
	OrderStatus_ORDERED    = OrderStatus("ORDERED")
	OrderStatus_MODIFYING  = OrderStatus("MODIFYING")
	OrderStatus_CANCELLING = OrderStatus("CANCELLING")
	OrderStatus_CANCELED   = OrderStatus("CANCELED")
	OrderStatus_EXECUTED   = OrderStatus("EXECUTED")
	OrderStatus_EXPIRED    = OrderStatus("EXPIRED")
)

type TimeInForce string

const (
	TimeInForce_FAK = TimeInForce("FAK")
	TimeInForce_FAS = TimeInForce("FAS")
)
