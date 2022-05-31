package basic

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go/api"
)

type Client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) *Client {
	return &Client{
		apiClient: apiClient,
	}
}

type LastBlock struct {
	Number int64
	Hash   string
}

func (c *Client) GetLastBlock() (*LastBlock, error) {

	endpoint := "block/number/latest"

	u, err := c.apiClient.MakeURL(endpoint, nil, api.Pagination{}, api.BasicFilters{})

	resp, err := c.apiClient.Get(u.String())

	var lastBlock LastBlock
	err = json.Unmarshal(resp.Data, &lastBlock)

	return &lastBlock, err
}

func (c *Client) GetBlockDetail() {

}

func (c *Client) GetAddressGasFee()    {}
func (c *Client) GetContractEvents()   {}
func (c *Client) RunContractFunction() {}
