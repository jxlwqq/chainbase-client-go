package basic

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go/api"
	"math/big"
	"time"
)

type Client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) *Client {
	return &Client{
		apiClient: apiClient,
	}
}

type LastBlock struct {
	Number int64
	Hash   string
}

func (c *Client) GetLastBlock() (*LastBlock, error) {

	endpoint := "block/number/latest"

	u, err := c.apiClient.MakeURL(endpoint, nil, nil, nil)

	resp, err := c.apiClient.Get(u.String())

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

func (c *Client) GetBlockDetail(number big.Int) (*BlockDetail, error) {

	endpoint := "block/detail"

	params := make(map[string]string)
	params["number"] = number.String()

	u, err := c.apiClient.MakeURL(endpoint, params, nil, nil)

	resp, err := c.apiClient.Get(u.String())

	var blockDetail BlockDetail
	err = json.Unmarshal(resp.Data, &blockDetail)

	return &blockDetail, err
}

type AddressGasFee struct {
	BlockNumber     big.Int `json:"block_number"`
	TransactionHash string  `json:"transaction_hash"`
	Fee             uint64  `json:"fee"`
}

func (c *Client) GetAddressGasFee(address string, pagination *api.Pagination, filters *api.BasicFilters) ([]*AddressGasFee, int, error) {

	endpoint := "account/fees/history"

	params := make(map[string]string)
	params["address"] = address

	u, err := c.apiClient.MakeURL(endpoint, params, pagination, filters)

	resp, err := c.apiClient.Get(u.String())

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

func (c *Client) GetContractEvents(contractAddress string, pagination *api.Pagination, filters *api.BasicFilters) ([]*ContractEvent, int, error) {

	endpoint := "contract/events"

	params := make(map[string]string)
	params["contract_address"] = contractAddress

	u, err := c.apiClient.MakeURL(endpoint, params, pagination, filters)

	resp, err := c.apiClient.Get(u.String())

	var contractEvents []*ContractEvent

	err = json.Unmarshal(resp.Data, &contractEvents)

	return contractEvents, resp.NextPage, err

}
func (c *Client) RunContractFunction() {}
