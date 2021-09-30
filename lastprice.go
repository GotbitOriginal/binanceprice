package binanceprice

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type LastPriceResp struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Price string `json:"price"`
}

func GetLastPrice(symbol string) (float64, error) {
	symbol = strings.Split(symbol, "_")[1] + strings.Split(symbol, "_")[0]
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + symbol

	data, err := MakeGetRequest(url)
	if err != nil {
		return 0.0, fmt.Errorf("error in GET price from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	var respApi LastPriceResp
	err = json.Unmarshal(data, &respApi)
	if err != nil {
		return 0.0, fmt.Errorf("error in unmarshl price from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	if respApi.Msg != "" || respApi.Code != 0 {
		return 0.0, fmt.Errorf("error with Binance api, sd: {%v}, rd: {%s}", url, string(data))
	}

	price, err := strconv.ParseFloat(respApi.Price, 64)
	if err != nil {
		return 0.0, fmt.Errorf("error in parse price Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	return price, nil
}
