package eth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/rahulrumalla/go-infura-json-rpc/infura"
)

type Client struct {
	*infura.Client
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

func NewClient(nw Network) *Client {
	return &Client{
		infura.NewClient(),
		NewConfig(os.Getenv("INFURA_API_KEY"), nw),
		"go-infura-json-rpc/0.1",
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

type Response struct {
	JSONRPCVersion string `json:"jsonrpc"`
	ID             uint8  `json:"id"`
	Result         string `json:"result"`
}
