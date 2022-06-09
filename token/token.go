package token

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go/web3api"
	"time"
)

type Client interface {
	GetBalance(address string) (string, error)
	GetTokenMetadata(contractAddress string) (*Metadata, error)
	GetAccountTokens(address string) ([]*AccountToken, error)
	GetAccountTransactions(address string) ([]*AccountTransaction, int, error)
	GetTokenHolders(contractAddress string) ([]string, error)
	GetTokenPrice(contractAddress string) (*Price, error)
	GetTokenPriceHistory(contractAddress string) ([]*Price, error)
}
type client struct {
	web3APIClient *web3api.Client
}

func New(web3APIClient *web3api.Client) Client {
	return &client{
		web3APIClient: web3APIClient,
	}
}

type Metadata struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	TotalSupply     string `json:"total_supply"`
}

type AccountToken struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	Balance         string `json:"balance"`
}

type AccountTransaction struct {
	Type                 int       `json:"type"`
	Status               int       `json:"status"`
	BlockNumber          int       `json:"block_number"`
	BlockTimestamp       time.Time `json:"block_timestamp"`
	TransactionHash      string    `json:"transaction_hash"`
	TransactionIndex     int       `json:"transaction_index"`
	FromAddress          string    `json:"from_address"`
	ToAddress            string    `json:"to_address"`
	Value                string    `json:"value"`
	Input                string    `json:"input"`
	Nonce                int       `json:"nonce"`
	ContractAddress      string    `json:"contract_address"`
	Gas                  int       `json:"gas"`
	GasPrice             int64     `json:"gas_price"`
	GasUsed              int       `json:"gas_used"`
	EffectiveGasPrice    int64     `json:"effective_gas_price"`
	CumulativeGasUsed    int       `json:"cumulative_gas_used"`
	MaxFeePerGas         int       `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas int       `json:"max_priority_fee_per_gas"`
	ChainId              int       `json:"chain_id"`
}

func (c *client) GetBalance(address string) (string, error) {
	endpoint := "account/balance"

	params := make(map[string]string)
	params["address"] = address

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	if err != nil {
		return "", err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return "", err
	}

	var balance string
	err = json.Unmarshal(resp.Data, &balance)

	return balance, nil

}

func (c *client) GetTokenMetadata(contractAddress string) (*Metadata, error) {
	endpoint := "token/metadata"

	params := make(map[string]string)
	params["contract_address"] = contractAddress

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	err = json.Unmarshal(resp.Data, &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (c *client) GetAccountTokens(address string) ([]*AccountToken, error) {

	endpoint := "account/tokens"

	params := make(map[string]string)
	params["address"] = address

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	var accountTokens []*AccountToken
	err = json.Unmarshal(resp.Data, &accountTokens)

	if err != nil {
		return nil, err
	}

	return accountTokens, nil

}

func (c *client) GetAccountTransactions(address string) ([]*AccountTransaction, int, error) {

	endpoint := "account/txs"

	params := make(map[string]string)
	params["address"] = address

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, 0, err
	}

	var accountTransactions []*AccountTransaction
	err = json.Unmarshal(resp.Data, &accountTransactions)

	if err != nil {
		return nil, 0, err
	}

	return accountTransactions, resp.NextPage, nil

}

func (c *client) GetTokenHolders(contractAddress string) ([]string, error) {

	endpoint := "token/holders"

	params := make(map[string]string)
	params["contract_address"] = contractAddress

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err

	}

	var addrs []string
	err = json.Unmarshal(resp.Data, &addrs)

	if err != nil {
		return nil, err
	}

	return addrs, nil
}

type Price struct {
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	Source    string    `json:"source"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *client) GetTokenPrice(contractAddress string) (*Price, error) {

	endpoint := "token/price"

	params := make(map[string]string)
	params["contract_address"] = contractAddress

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	var price Price

	err = json.Unmarshal(resp.Data, &price)

	if err != nil {
		return nil, err
	}

	return &price, nil
}

func (c *client) GetTokenPriceHistory(contractAddress string) ([]*Price, error) {

	endpoint := "token/price/history"

	params := make(map[string]string)
	params["contract_address"] = contractAddress

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	var prices []*Price

	err = json.Unmarshal(resp.Data, &prices)

	if err != nil {
		return nil, err
	}

	return prices, nil

}
