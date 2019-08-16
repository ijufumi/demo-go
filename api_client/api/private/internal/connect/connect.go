package connect

import (
	"api_client/api/common/configuration"
	"api_client/api/private/internal/auth"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const host = "https://api.coin.z.com/private"

// Connection ...
type Connection struct {
	apiKey    string
	secretKey string
}

// New ...
func New(apiKey, secretKey string) *Connection {
	return &Connection{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// Post ...
func (c *Connection) Post(body interface{}, path string) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqeuest, err := http.NewRequest("POST", host+path, strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	c.makeHeader(reqeuest, "POST", path, string(b))

	if configuration.Debug {
		fmt.Printf("[Request]Header:%v\n", reqeuest.Header)
		fmt.Printf("[Request]Body:%v\n", string(b))
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

// Get ...
func (c *Connection) Get(param url.Values, path string) ([]byte, error) {
	queryString := param.Encode()
	reqeuest, err := http.NewRequest("GET", host+path+"?"+queryString, nil)
	if err != nil {
		return nil, err
	}
	c.makeHeader(reqeuest, "GET", path, "")

	if configuration.Debug {
		fmt.Printf("[Request]Header:%v\n", reqeuest.Header)
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

func (c *Connection) makeHeader(r *http.Request, method, path, body string) {
	timeStamp := time.Now().Unix() * 1000
	r.Header.Set("API-TIMESTAMP", fmt.Sprint(timeStamp))
	r.Header.Set("API-KEY", c.apiKey)
	r.Header.Set("API-SIGN", auth.MakeSign(c.secretKey, timeStamp, method, path, body))
}
