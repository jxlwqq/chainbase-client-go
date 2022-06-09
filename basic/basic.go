package basic

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go/web3api"
	"math/big"
	"time"
)

type Client interface {
	GetLastBlock() (*LastBlock, error)
	GetBlockDetail(number big.Int) (*BlockDetail, error)
	GetAddressGasFee(address string, pagination *web3api.Pagination, filters *web3api.BasicFilters) ([]*AddressGasFee, int, error)
	GetContractEvents(contractAddress string, pagination *web3api.Pagination, filters *web3api.BasicFilters) ([]*ContractEvent, int, error)
	RunContractFunction()
}

type client struct {
	web3APIClient *web3api.Client
}

func New(web3APIClient *web3api.Client) Client {
	return &client{
		web3APIClient: web3APIClient,
	}
}

type LastBlock struct {
	Number big.Int
	Hash   string
}

func (c *client) GetLastBlock() (*LastBlock, error) {

	endpoint := "block/number/latest"

	url, err := c.web3APIClient.MakeURL(endpoint, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())

	if err != nil {
		return nil, err
	}

	var lastBlock LastBlock
	err = json.Unmarshal(resp.Data, &lastBlock)

	return &lastBlock, err
}

type BlockDetail struct {
	Number            big.Int   `json:"number"`
	Hash              string    `json:"hash"`
	ParentHash        string    `json:"parent_hash"`
	Timestamp         time.Time `json:"timestamp"`
	Difficulty        string    `json:"difficulty"`
	TotalDifficulty   string    `json:"total_difficulty"`
	ExtraData         string    `json:"extra_data"`
	GasLimit          uint64    `json:"gas_limit"`
	GasUsed           uint64    `json:"gas_used"`
	BaseFeePerGas     uint64    `json:"base_fee_per_gas"`
	LogsBloom         string    `json:"logs_bloom"`
	Miner             string    `json:"miner"`
	MixHash           string    `json:"mix_hash"`
	Nonce             string    `json:"nonce"`
	ReceiptsRoot      string    `json:"receipts_root"`
	Sha3Uncles        string    `json:"sha3_uncles"`
	Size              int       `json:"size"`
	StateRoot         string    `json:"state_root"`
	TransactionsRoot  string    `json:"transactions_root"`
	TransactionsCount int       `json:"transactions_count"`
	UnclesCount       int       `json:"uncles_count"`
	ChainId           int       `json:"chain_id"`
}

func (c *client) GetBlockDetail(number big.Int) (*BlockDetail, error) {

	endpoint := "block/detail"

	params := make(map[string]string)
	params["number"] = number.String()

	url, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.web3APIClient.Get(url.String())

	if err != nil {
		return nil, err
	}

	var blockDetail BlockDetail
	err = json.Unmarshal(resp.Data, &blockDetail)

	return &blockDetail, err
}

type AddressGasFee struct {
	BlockNumber     big.Int `json:"block_number"`
	TransactionHash string  `json:"transaction_hash"`
	Fee             uint64  `json:"fee"`
}

func (c *client) GetAddressGasFee(address string, pagination *web3api.Pagination, filters *web3api.BasicFilters) ([]*AddressGasFee, int, error) {

	endpoint := "account/fees/history"

	params := make(map[string]string)
	params["address"] = address

	url, err := c.web3APIClient.MakeURL(endpoint, params, pagination, filters)

	if err != nil {
		return nil, 0, err
	}

	resp, err := c.web3APIClient.Get(url.String())

	if err != nil {
		return nil, 0, err
	}

	var addressGasFee []*AddressGasFee
	err = json.Unmarshal(resp.Data, &addressGasFee)

	return addressGasFee, resp.NextPage, err
}

type ContractEvent struct {
	BlockNumber      big.Int `json:"block_number"`
	TransactionHash  string  `json:"transaction_hash"`
	TransactionIndex int     `json:"transaction_index"`
	MethodID         string  `json:"method_id"`
	Function         string  `json:"function"`
}

func (c *client) GetContractEvents(contractAddress string, pagination *web3api.Pagination, filters *web3api.BasicFilters) ([]*ContractEvent, int, error) {

	endpoint := "contract/events"

	params := make(map[string]string)
	params["contract_address"] = contractAddress

	url, err := c.web3APIClient.MakeURL(endpoint, params, pagination, filters)

	if err != nil {
		return nil, 0, err
	}

	resp, err := c.web3APIClient.Get(url.String())

	if err != nil {
		return nil, 0, err
	}

	var contractEvents []*ContractEvent

	err = json.Unmarshal(resp.Data, &contractEvents)

	return contractEvents, resp.NextPage, err

}
func (c *client) RunContractFunction() {
	_ = "contract/call"
}
