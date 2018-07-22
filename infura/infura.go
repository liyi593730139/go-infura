package infura

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const (
	baseURL = ""
)

type Client struct {
	*http.Client
	config    Config
	userAgent string
}

type Request struct {
	Endpoint       string   `json:"method"`
	HTTPMethod     string   `json:"omit"`
	Params         []string `json:"params"`
	JSONRPCVersion string   `json:"jsonrpc"`
	ID             uint8    `json:"id"`
}

func NewClient(nw Network) Client {
	return Client{
		http.DefaultClient,
		NewConfig(os.Getenv("INFURA_API_KEY"), nw),
		"go-infure-json-rpc/0.1",
	}
}

func (c Client) NewRequest(httpMethod, endpoint string, params []string) (*http.Request, error) {
	req := Request{
		JSONRPCVersion: "2.0",
		ID:             1,
		Endpoint:       endpoint,
		Params:         params,
	}

	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(httpMethod, c.config.Network.URL()+c.config.APIKey, buf)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Accept", "application/json")
	if c.userAgent != "" {
		httpReq.Header.Add("User-Agent", c.userAgent)
	}

	return httpReq, err
}

func (c Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}

type Response struct {
	JSONRPCVersion string `json:"jsonrpc"`
	ID             uint8  `json:"id"`
	Result         string `json:"result"`
}
