package chainnetwork

import (
	"encoding/json"
	"fmt"
	"github.com/jxlwqq/chainbase-client-go/chain"
	"io"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://%s.s.chainbase.online/v1/%s"

type Client struct {
	httpClient *http.Client
	chainID    chain.ID
	apiKey     string
}

type Response struct {
	JSONRPC string          `json:"json_rpc"`
	ID      string          `json:"id"`
	Result  json.RawMessage `json:"result"`
}

var ChainIDSLDPrefixMap = map[chain.ID]string{
	chain.EthereumMainnet: "ethereum-mainnet",
	chain.EthereumRinkeby: "ethereum-rinkeby",
	chain.PolygonMainnet:  "polygon-mainnet",
	chain.PolygonMumbai:   "polygon-mumbai",
	chain.BSCMainnet:      "bsc-mainnet",
	chain.BSCTestnet:      "bsc-testnet",
}

func New(httpClient *http.Client, chainID chain.ID, apiKey string) *Client {

	return &Client{
		httpClient: httpClient,
		chainID:    chainID,
		apiKey:     apiKey,
	}
}

func (c *Client) MakeURL() (*url.URL, error) {
	sldPrefix := ChainIDSLDPrefixMap[c.chainID]

	u, err := url.Parse(fmt.Sprintf(defaultBaseURL, sldPrefix, c.apiKey))

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (c *Client) Post(url string, body io.Reader) (*Response, error) {
	return c.Do(http.MethodPost, url, body)
}

func (c *Client) Do(method string, url string, body io.Reader) (*Response, error) {
	req, _ := http.NewRequest(method, url, body)

	req.Header.Set("Content-Type", "application/json;")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &response, err
}
