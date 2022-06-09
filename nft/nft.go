package nft

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go/web3api"
	"math/big"
	"time"
)

type Client interface {
	SearchNFTs(name string, contractAddress string) ([]*NFT, error)
	GetAccountNFTs(address string) ([]*AccountNFT, error)
	GetNFTTransfers(contractAddress string, tokenID *big.Int) ([]*Transfer, error)
	GetNFTMetadata(contractAddress string, tokenID *big.Int) (*Metadata, error)
	GetNFTOwner(contractAddress string, tokenID *big.Int) (string, error)
	GetNFTOwnerHistory(contractAddress string, tokenID *big.Int) ([]*OwnerHistory, error)
	GetNFTFloorPrice(contractAddress string) (*FloorPrice, error)
}
type client struct {
	web3APIClient *web3api.Client
}

func New(web3APIClient *web3api.Client) Client {
	return &client{
		web3APIClient: web3APIClient,
	}
}

type NFT struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Owner           string `json:"owner"`
}

func (c *client) SearchNFTs(name string, contractAddress string) ([]*NFT, error) {
	endpoint := "nft/search"
	params := make(map[string]string)
	params["name"] = name
	if contractAddress != "" {
		params["contract_address"] = contractAddress
	}

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	var nfts []*NFT
	err = json.Unmarshal(resp.Data, &nfts)
	if err != nil {
		return nil, err
	}

	return nfts, nil
}

type AccountNFT struct {
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
	ContractType    string `json:"contract_type"`
	ContractName    string `json:"contract_name"`
	ContractSymbol  string `json:"contract_symbol"`
	TokenURI        string `json:"token_uri"`
}

func (c *client) GetAccountNFTs(address string) ([]*AccountNFT, error) {
	endpoint := "account/nfts"
	params := make(map[string]string)
	params["address"] = address

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())

	var accountNFTs []*AccountNFT

	err = json.Unmarshal(resp.Data, &accountNFTs)

	if err != nil {
		return nil, err
	}

	return accountNFTs, nil
}

type Transfer struct {
	BlockTimestamp   time.Time `json:"block_timestamp"`
	BlockNumber      int       `json:"block_number"`
	TransactionHash  string    `json:"transaction_hash"`
	TransactionIndex int       `json:"transaction_index"`
	FromAddress      string    `json:"from_address"`
	ToAddress        string    `json:"to_address"`
	Value            string    `json:"value"`
	TokenId          string    `json:"token_id"`
	OperatorAddress  string    `json:"operator_address"`
	LogIndex         int       `json:"log_index"`
	ChainId          int       `json:"chain_id"`
}

func (c *client) GetNFTTransfers(contractAddress string, tokenID *big.Int) ([]*Transfer, error) {
	endpoint := "nft/transfers"
	params := make(map[string]string)
	params["contract_address"] = contractAddress
	if tokenID.BitLen() != 0 {
		params["token_id"] = tokenID.String()
	}

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())

	if err != nil {
		return nil, err
	}

	var transfers []*Transfer
	err = json.Unmarshal(resp.Data, &transfers)

	if err != nil {
		return nil, err
	}

	return transfers, nil
}

type Metadata struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        int    `json:"decimals"`
	Owner           string `json:"owner"`
	TokenUri        string `json:"token_uri"`
	TokenID         string `json:"token_id"`
}

func (c *client) GetNFTMetadata(contractAddress string, tokenID *big.Int) (*Metadata, error) {
	endpoint := "nft/metadata"

	params := make(map[string]string)
	params["contract_address"] = contractAddress
	params["token_id"] = tokenID.String()

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

	return &metadata, nil
}

func (c *client) GetNFTOwner(contractAddress string, tokenID *big.Int) (string, error) {
	endpoint := "nft/owner"

	params := make(map[string]string)
	params["contract_address"] = contractAddress
	params["token_id"] = tokenID.String()

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return "", err
	}

	var addr string

	err = json.Unmarshal(resp.Data, &addr)
	if err != nil {
		return "", err
	}
	return addr, nil

}

type OwnerHistory struct {
	BlockTimestamp   time.Time `json:"block_timestamp"`
	BlockNumber      int       `json:"block_number"`
	TransactionHash  string    `json:"transaction_hash"`
	TransactionIndex int       `json:"transaction_index"`
	FromAddress      string    `json:"from_address"`
	ToAddress        string    `json:"to_address"`
	Value            string    `json:"value"`
	TokenId          string    `json:"token_id"`
	OperatorAddress  string    `json:"operator_address"`
	LogIndex         int       `json:"log_index"`
	ChainId          int       `json:"chain_id"`
}

func (c *client) GetNFTOwnerHistory(contractAddress string, tokenID *big.Int) ([]*OwnerHistory, error) {
	endpoint := "nft/owner/history"

	params := make(map[string]string)
	params["contract_address"] = contractAddress
	params["token_id"] = tokenID.String()

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	var ownerHistories []*OwnerHistory

	err = json.Unmarshal(resp.Data, &ownerHistories)
	if err != nil {
		return nil, err
	}
	return ownerHistories, nil
}

type FloorPrice struct {
	FloorPrice float64   `json:"floor_price"`
	Symbol     string    `json:"symbol"`
	Source     string    `json:"source"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c *client) GetNFTFloorPrice(contractAddress string) (*FloorPrice, error) {
	endpoint := "nft/floor_price"

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

	var floorPrice FloorPrice

	err = json.Unmarshal(resp.Data, &floorPrice)
	if err != nil {
		return nil, err
	}
	return &floorPrice, nil
}
