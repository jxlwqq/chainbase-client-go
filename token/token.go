package token

import "github.com/jxlwqq/chainbase-client-go/api"

type Client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) *Client {
	return &Client{apiClient: apiClient}
}

func (c *Client) GetBalance() {
}

func (c *Client) GetTokenBalance() {
}

func (c *Client) GetAccountTokens() {
}

func (c *Client) GetAccountTransactions() {
}

func (c *Client) GetTokenHolders() {
}

func (c *Client) GetTokenPrice() {
}

func (c *Client) GetTokenPriceHistory() {
}
