// Code generated by "enumer -output=enum_string.go -type ExchangeType -trimprefix ExchangeType"; DO NOT EDIT.

package apis

import (
	"fmt"
	"strings"
)

const _ExchangeTypeName = "BinanceHuobiOkexBW"

var _ExchangeTypeIndex = [...]uint8{0, 7, 12, 16, 18}

const _ExchangeTypeLowerName = "binancehuobiokexbw"

func (i ExchangeType) String() string {
	i -= 1
	if i >= ExchangeType(len(_ExchangeTypeIndex)-1) {
		return fmt.Sprintf("ExchangeType(%d)", i+1)
	}
	return _ExchangeTypeName[_ExchangeTypeIndex[i]:_ExchangeTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ExchangeTypeNoOp() {
	var x [1]struct{}
	_ = x[ExchangeTypeBinance-(1)]
	_ = x[ExchangeTypeHuobi-(2)]
	_ = x[ExchangeTypeOkex-(3)]
	_ = x[ExchangeTypeBW-(4)]
}

var _ExchangeTypeValues = []ExchangeType{ExchangeTypeBinance, ExchangeTypeHuobi, ExchangeTypeOkex, ExchangeTypeBW}

var _ExchangeTypeNameToValueMap = map[string]ExchangeType{
	_ExchangeTypeName[0:7]:        ExchangeTypeBinance,
	_ExchangeTypeLowerName[0:7]:   ExchangeTypeBinance,
	_ExchangeTypeName[7:12]:       ExchangeTypeHuobi,
	_ExchangeTypeLowerName[7:12]:  ExchangeTypeHuobi,
	_ExchangeTypeName[12:16]:      ExchangeTypeOkex,
	_ExchangeTypeLowerName[12:16]: ExchangeTypeOkex,
	_ExchangeTypeName[16:18]:      ExchangeTypeBW,
	_ExchangeTypeLowerName[16:18]: ExchangeTypeBW,
}

var _ExchangeTypeNames = []string{
	_ExchangeTypeName[0:7],
	_ExchangeTypeName[7:12],
	_ExchangeTypeName[12:16],
	_ExchangeTypeName[16:18],
}

// ExchangeTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ExchangeTypeString(s string) (ExchangeType, error) {
	if val, ok := _ExchangeTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ExchangeTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ExchangeType values", s)
}

// ExchangeTypeValues returns all values of the enum
func ExchangeTypeValues() []ExchangeType {
	return _ExchangeTypeValues
}

// ExchangeTypeStrings returns a slice of all String values of the enum
func ExchangeTypeStrings() []string {
	strs := make([]string, len(_ExchangeTypeNames))
	copy(strs, _ExchangeTypeNames)
	return strs
}

// IsAExchangeType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ExchangeType) IsAExchangeType() bool {
	for _, v := range _ExchangeTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
