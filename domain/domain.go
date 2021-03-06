package domain

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go/web3api"
	"time"
)

type Client interface {
	GetENSRecords(domain string) (*ENSRecord, error)
	GetENSReverse(address string) ([]ENSRecord, error)
	GetAccountENS(address string) ([]ENSRecord, error)
}

type client struct {
	web3APIClient *web3api.Client
}

func New(web3APIClient *web3api.Client) Client {
	return &client{
		web3APIClient: web3APIClient,
	}
}

type ENSRecord struct {
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	Registrant     string    `json:"registrant"`
	Owner          string    `json:"owner"`
	Resolver       string    `json:"resolver"`
	RegistrantTime time.Time `json:"registrant_time"`
	ExpirationTime time.Time `json:"expiration_time"`
	TokenId        string    `json:"token_id"`
}

func (c *client) GetENSRecords(domain string) (*ENSRecord, error) {
	endpoint := "ens/records"
	params := make(map[string]string)
	params["domain"] = domain
	u, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	resp, err := c.web3APIClient.Get(u.String())

	var ensRecord ENSRecord
	err = json.Unmarshal(resp.Data, &ensRecord)

	return &ensRecord, err

}

func (c *client) GetENSReverse(address string) ([]ENSRecord, error) {

	endpoint := "ens/reverse"
	params := make(map[string]string)
	params["address"] = address
	u, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	resp, err := c.web3APIClient.Get(u.String())

	var ensRecords []ENSRecord

	err = json.Unmarshal(resp.Data, &ensRecords)

	return ensRecords, err
}

func (c *client) GetAccountENS(address string) ([]ENSRecord, error) {

	endpoint := "account/ens"
	params := make(map[string]string)
	params["address"] = address
	u, err := c.web3APIClient.MakeURL(endpoint, params, nil, nil)

	resp, err := c.web3APIClient.Get(u.String())

	var ensRecords []ENSRecord

	err = json.Unmarshal(resp.Data, &ensRecords)

	return ensRecords, err
}
