package ipfs

const (
	getObjectPath = "object/get"
	nodeInfoPath  = "id"
)

type ObjectLink struct {
	Name string
	Hash string
	Size uint64
}

type Object struct {
	Links []ObjectLink
	Data  string
}

type NodeInfo struct {
	ID              string
	PublicKey       string
	Addresses       []string
	AgentVersion    string
	ProtocolVersion string
}

func (c *Client) GetObject(key string) (*Object, error) {
	params := make(map[string]string)
	params["arg"] = key

	req, err := c.NewRequest(getObjectPath, params)
	if err != nil {
		return nil, err
	}

	var resp = new(Object)
	_, err = c.Call(req, resp)

	return resp, err
}

func (c *Client) NodeInfo() (*NodeInfo, error) {
	req, err := c.NewRequest(nodeInfoPath, nil)
	if err != nil {
		return nil, err
	}

	var resp = new(NodeInfo)
	_, err = c.Call(req, resp)

	return resp, err
}
