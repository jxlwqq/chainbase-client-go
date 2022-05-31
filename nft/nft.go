package nft

import "github.com/jxlwqq/chainbase-client-go/api"

type Client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) *Client {
	return &Client{apiClient}
}

func (c *Client) SearchNFTs() {

}

func (c *Client) GetAccountNFTs() {
}

func (c *Client) GetNFTTransfers() {

}

func (c *Client) GetNFTMetadata() {

}

func (c *Client) GetNFTOwner() {

}

func (c *Client) GetNFTOwnerHistory() {

}

func (c *Client) GetNFTFloorPrice() {

}
