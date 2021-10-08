package binanceprice

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func GetVolumePeak(symbol string) (bool, error) {
	symbol = strings.Split(symbol, "_")[1] + strings.Split(symbol, "_")[0]
	url := "https://api.binance.com/api/v3/klines?symbol=" + symbol + "&interval=5m&limit=3"

	data, err := MakeGetRequest(url)
	if err != nil {
		return false, fmt.Errorf("error in GET kline from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	var respApi [][]interface{}
	err = json.Unmarshal(data, &respApi)
	if err != nil {
		return false, fmt.Errorf("error in unmarshl kline from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	if len(respApi) != 3 {
		return false, fmt.Errorf("error with kline response from Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	// 5 - volume
	lastVolumeString, ok := respApi[1][5].(string)
	if !ok {
		return false, fmt.Errorf("error with parse last volume in string from kline Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}
	lastVolume, err := strconv.ParseFloat(lastVolumeString, 64)
	if !ok {
		return false, fmt.Errorf("error with parse last volume in float64 from kline Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	previousVolumeString, ok := respApi[0][5].(string)
	if !ok {
		return false, fmt.Errorf("error with parse previous volume in string from kline Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}
	previousVolume, err := strconv.ParseFloat(previousVolumeString, 64)
	if !ok {
		return false, fmt.Errorf("error with parse previous volume in float64 from kline Binance, sd: {%v}, rd: {%s}, err: %s", url, string(data), err.Error())
	}

	if (lastVolume/previousVolume >= 2.0 && lastVolume >= 1000.0) || (previousVolume >= 1200.0 && lastVolume >= 1200.0) {
		return true, nil
	} else {
		return false, nil
	}
}
