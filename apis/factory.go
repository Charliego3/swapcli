package apis

type Exchange interface {
	CancelOrders() error
	CreateOrder()
}

func Fetch(t ExchangeType) Exchange {
	switch t {
	case ExchangeTypeBinance:
		return &Binance{}
	}
	return nil
}
