package chainnetwork

import (
	"github.com/jxlwqq/chainbase-client-go/chain"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	chainID    chain.ID
	apiKey     string
}

func New(httpClient *http.Client, chainID chain.ID, apiKey string) *Client {
	return &Client{
		httpClient: httpClient,
		chainID:    chainID,
		apiKey:     apiKey,
	}
}
