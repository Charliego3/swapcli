package apis

type Binance struct {
}

func (b *Binance) SetSigner(signer Signer) {

}

func (b *Binance) Signer() Signer {
	return nil
}

func (b *Binance) CancelOrders() error {
	return nil
}

func (b *Binance) CreateOrder() {

}
