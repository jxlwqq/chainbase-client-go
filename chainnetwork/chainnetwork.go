package chainnetwork

type Client struct {
	APIKey string
}

func New(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}
