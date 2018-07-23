package eth

type Config struct {
	APIKey  string
	Network Network
}

func NewConfig(key string, nw Network) Config {
	return Config{
		APIKey:  key,
		Network: nw,
	}
}
