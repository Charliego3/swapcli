package apis

type Signer interface {
	Sign() (string, error)
	CheckAccount()
}

type Exchange interface {
	SetSigner(Signer)
	Signer() Signer
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
