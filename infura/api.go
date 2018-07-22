package infura

import "net/http"

const (
	getAccountBalance = "eth_getBalance"
	getBlockNumber    = "eth_blockNumber"
)

type BlockParameter string

var (
	Latest   BlockParameter = "latest"
	Earliest BlockParameter = "earliest"
	Pending  BlockParameter = "pending"
)

func (c *Client) GetAccountBalance(address string, blockParameter BlockParameter) (string, error) {
	params := []string{address, string(blockParameter)}
	req, err := c.NewRequest(http.MethodPost, getAccountBalance, params)
	if err != nil {
		return "", err
	}

	var resp = new(Response)
	_, err = c.do(req, resp)

	return resp.Result, err
}

func (c *Client) GetBlockNumber() (string, error) {
	req, err := c.NewRequest(http.MethodPost, getBlockNumber, []string{})
	if err != nil {
		return "", err
	}

	var resp = new(Response)
	_, err = c.do(req, resp)

	return resp.Result, err
}
