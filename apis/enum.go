package apis

//go:generate enumer -output=enum_string.go -type ExchangeType -trimprefix ExchangeType
type ExchangeType uint

const (
	ExchangeTypeBinance ExchangeType = iota + 1
	ExchangeTypeHuobi
	ExchangeTypeOkex
	ExchangeTypeBW
)
