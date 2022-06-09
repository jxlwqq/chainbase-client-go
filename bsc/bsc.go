package bsc

import (
	"github.com/jxlwqq/chainbase-client-go/chainnetwork"
)

type Client interface {
}
type client struct {
	chainNetworkClient *chainnetwork.Client
}

func New(chainNetworkClient *chainnetwork.Client) Client {
	return &client{
		chainNetworkClient: chainNetworkClient,
	}
}
