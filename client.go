package chainbase

import (
	"github.com/jxlwqq/chainbase-client-go/api"
	"github.com/jxlwqq/chainbase-client-go/basic"
	"github.com/jxlwqq/chainbase-client-go/domain"
	"github.com/jxlwqq/chainbase-client-go/token"
	"net/http"
)

type Client struct {
	Basic  *basic.Client
	Domain *domain.Client
	Token  *token.Client
}

func New(httpClient *http.Client, chainID int, apiKey string) *Client {
	apiClient := api.New(httpClient, chainID, apiKey)
	return &Client{
		Basic:  basic.New(apiClient),
		Domain: domain.New(apiClient),
		Token:  token.New(apiClient),
	}
}
