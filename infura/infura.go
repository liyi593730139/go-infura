package infura

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		http.DefaultClient,
	}
}

func (c *Client) Call(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	// response, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(string(response))

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

func BuildURLString(u *url.URL, params map[string]string) string {
	u.RawQuery = encodeParams(u, params)
	return u.String()
}

func encodeParams(u *url.URL, params map[string]string) string {
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	return q.Encode()
}
