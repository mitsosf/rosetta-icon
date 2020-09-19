package helpers

import (
	"strconv"
	"strings"
)

func strip0xFromHexString(hexValue string) string {
	// replace 0x or 0X with empty String
	return strings.Replace(hexValue, "0x", "x", -1)
}

func HexToDecimal(hexValue string) string{
	//TODO FIX big int issue
	result, err := strconv.ParseInt(strip0xFromHexString(hexValue), 16, 64)

	if err != nil {
		panic(err)
	}

	return strconv.FormatInt(result, 10)
}
