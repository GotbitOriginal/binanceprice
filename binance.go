package binanceprice

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type BinanceResp struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Price string `json:"price"`
}

func GetLastPriceBinance(symbol string) (float64, error) {
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + symbol

	data, err := MakeGetRequest(url)
	if err != nil {
		return 0.0, fmt.Errorf("Error in GET price from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	var respApi BinanceResp
	err = json.Unmarshal(data, &respApi)
	if err != nil {
		return 0.0, fmt.Errorf("Error in unmarshl price from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	if respApi.Msg != "" || respApi.Code != 0 {
		return 0.0, fmt.Errorf("Error with Binance api, sd: {%v}, rd: {%s}", url, string(data))
	}

	price, err := strconv.ParseFloat(respApi.Price, 64)
	if err != nil {
		return 0.0, fmt.Errorf("Error in parse price Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	return price, nil
}
