package helpers

import (
	"github.com/ethereum/go-ethereum/common/math"
)

func HexToDecimal(hexValue string) string {
	result, ok := math.ParseBig256(hexValue)

	if !ok {
		panic("invalid hex or decimal integer")
	}

	return result.String()
}
