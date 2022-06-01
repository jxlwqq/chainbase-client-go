package polygon

import "github.com/jxlwqq/chainbase-client-go/api"

type Client interface {
}

type client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) Client {
	return &client{apiClient}
}
