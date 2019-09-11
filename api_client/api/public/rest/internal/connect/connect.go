package connect

import (
	"api_client/api/common/configuration"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const host = "https://api.coin.z.com/public"

// Connection ...
type Connection struct{}

// Get ...
func (c *Connection) Get(param url.Values, path string) ([]byte, error) {
	queryString := param.Encode()
	urlString := host + path
	if len(queryString) != 0 {
		urlString = urlString + "?" + queryString
	}
	reqeuest, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	if configuration.Debug {
		fmt.Printf("[Request]URL:%v\n", reqeuest.URL)
	}

	response, err := http.DefaultClient.Do(reqeuest)
	if err != nil {
		return nil, err
	}

	defer func() {
		response.Body.Close()
	}()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if configuration.Debug {
		fmt.Printf("[Response]Body:%v\n", string(resBody))
	}

	return resBody, nil

}
