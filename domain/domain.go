package domain

import "github.com/jxlwqq/chainbase-client-go/api"

type Client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) *Client {
	return &Client{
		apiClient: apiClient,
	}
}

func (c *Client) GetENSRecords() {

}

func (c *Client) GetENSReverse() {

}

func (c *Client) GetAccountENS() {

}
