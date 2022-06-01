package nft

import "github.com/jxlwqq/chainbase-client-go/api"

type Client interface {
	SearchNFTs()
	GetAccountNFTs()
	GetNFTTransfers()
	GetNFTMetadata()
	GetNFTOwner()
	GetNFTOwnerHistory()
	GetNFTFloorPrice()
}
type client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) Client {
	return &client{apiClient}
}

func (c *client) SearchNFTs() {

}

func (c *client) GetAccountNFTs() {
}

func (c *client) GetNFTTransfers() {

}

func (c *client) GetNFTMetadata() {

}

func (c *client) GetNFTOwner() {

}

func (c *client) GetNFTOwnerHistory() {

}

func (c *client) GetNFTFloorPrice() {

}
