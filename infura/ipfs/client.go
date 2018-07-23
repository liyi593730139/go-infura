package ipfs

import (
	"net/http"
	"net/url"

	"github.com/rahulrumalla/go-infura-json-rpc/infura"
)

const (
	baseAPIURL     = "https://ipfs.infura.io:5001/api/v0/"
	baseGatewayURL = "https://ipfs.infura.io/ipfs/"
)

type Client struct {
	*infura.Client
	userAgent string

	baseAPIURL     string
	baseAPIGateway string
}

func NewClient() *Client {
	return &Client{
		infura.NewClient(),
		"go-infura-ipfs/0.1",
		baseAPIURL,
		baseGatewayURL,
	}
}

func (c *Client) NewRequest(relPath string, params map[string]string) (*http.Request, error) {
	u, err := url.Parse(c.baseAPIURL + relPath)
	if err != nil {
		return nil, err
	}

	var url string
	if params != nil {
		url = infura.BuildURLString(u, params)
	} else {
		url = u.String()
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if c.userAgent != "" {
		req.Header.Add("User-Agent", c.userAgent)
	}

	return req, nil
}
