package chainbase

import (
	"github.com/jxlwqq/chainbase-client-go/basic"
	"github.com/jxlwqq/chainbase-client-go/bsc"
	"github.com/jxlwqq/chainbase-client-go/chain"
	"github.com/jxlwqq/chainbase-client-go/chainnetwork"
	"github.com/jxlwqq/chainbase-client-go/domain"
	"github.com/jxlwqq/chainbase-client-go/ethereum"
	"github.com/jxlwqq/chainbase-client-go/nft"
	"github.com/jxlwqq/chainbase-client-go/polygon"
	"github.com/jxlwqq/chainbase-client-go/token"
	"github.com/jxlwqq/chainbase-client-go/web3api"
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

func New(httpClient *http.Client, chainID chain.ID, web3APIKey string, chainNetworkAPIKey string) *Client {
	web3APIClient := web3api.New(httpClient, chainID, web3APIKey)
	chainNetworkClient := chainnetwork.New(httpClient, chainID, chainNetworkAPIKey)
	return &Client{
		Basic:    basic.New(web3APIClient),
		NFT:      nft.New(web3APIClient),
		Token:    token.New(web3APIClient),
		Domain:   domain.New(web3APIClient),
		Ethereum: ethereum.New(chainNetworkClient),
		Polygon:  polygon.New(chainNetworkClient),
		BSC:      bsc.New(chainNetworkClient),
	}
}
