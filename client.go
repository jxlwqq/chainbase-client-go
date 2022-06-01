package chainbase

import (
	"github.com/jxlwqq/chainbase-client-go/api"
	"github.com/jxlwqq/chainbase-client-go/basic"
	"github.com/jxlwqq/chainbase-client-go/bsc"
	"github.com/jxlwqq/chainbase-client-go/domain"
	"github.com/jxlwqq/chainbase-client-go/ethereum"
	"github.com/jxlwqq/chainbase-client-go/nft"
	"github.com/jxlwqq/chainbase-client-go/polygon"
	"github.com/jxlwqq/chainbase-client-go/token"
	"net/http"
)

type Client struct {
	Basic    basic.Client
	NFT      nft.Client
	Token    token.Client
	Domain   domain.Client
	Ethereum ethereum.Client
	Polygon  polygon.Client
	BSC      bsc.Client
}

func New(httpClient *http.Client, chainID api.ChainID, apiKey string) *Client {
	apiClient := api.New(httpClient, chainID, apiKey)
	return &Client{
		Basic:    basic.New(apiClient),
		NFT:      nft.New(apiClient),
		Token:    token.New(apiClient),
		Domain:   domain.New(apiClient),
		Ethereum: ethereum.New(apiClient),
		Polygon:  polygon.New(apiClient),
		BSC:      bsc.New(apiClient),
	}
}
