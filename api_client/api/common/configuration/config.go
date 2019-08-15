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
