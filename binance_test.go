package binanceprice

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	symbol := "USDT_BTC"

	// price, err := GetLastPrice(symbol)
	// fmt.Println(price, err)

	peac, err := GetVolumePeak(symbol)
	fmt.Println(peac, err)
}
