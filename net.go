package binance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	DEFAULT_TIMEOUT = 10 // S
)

func makeRequest(request *http.Request, timeout int) (result []byte, err error) {
	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	response, err := netClient.Do(request)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf(
			"server response with status: %v, with message: %v",
			response.StatusCode,
			string(result),
		)
	}

	return result, nil
}

func MakeGetRequest(url string) (result []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	return makeRequest(request, DEFAULT_TIMEOUT)
}
