package api

import (
	"encoding/json"
	"github.com/jxlwqq/chainbase-client-go"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const defaultBaseURL = "https://api.chainbase.online/v1/"

type Client struct {
	httpClient *http.Client
	ChainID    chainbase.ChainID
	APIKey     string
}

type Response struct {
	Code     int32           `json:"code"`
	Message  string          `json:"message"`
	Data     json.RawMessage `json:"data"`
	NextPage int             `json:"next_page"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type BasicFilters struct {
	FromTimestamp int64 `json:"from_timestamp"`
	EndTimestamp  int64 `json:"end_timestamp"`
	FromBlock     int64 `json:"from_block"`
	ToBlock       int64 `json:"to_block"`
}

func New(httpClient *http.Client, chainID chainbase.ChainID, apiKey string) *Client {
	return &Client{
		httpClient: httpClient,
		ChainID:    chainID,
		APIKey:     apiKey,
	}
}

func (c *Client) MakeURL(endpoint string, parameters map[string]string, pagination *Pagination, filters *BasicFilters) (*url.URL, error) {
	u, _ := url.Parse(defaultBaseURL + endpoint)

	q := u.Query()
	q.Set("chain_id", c.ChainID.String())

	for k, v := range parameters {
		q.Set(k, v)
	}

	if pagination.Page > 0 {
		q.Set("page", strconv.Itoa(pagination.Page))
	}

	if pagination.Limit > 0 {
		q.Set("limit", strconv.Itoa(pagination.Limit))
	}

	if filters.FromTimestamp > 0 {
		q.Set("from_timestamp", strconv.FormatInt(filters.FromTimestamp, 10))
	}

	if filters.EndTimestamp > 0 {
		q.Set("end_timestamp", strconv.FormatInt(filters.EndTimestamp, 10))
	}

	if filters.FromBlock > 0 {
		q.Set("from_block", strconv.FormatInt(filters.FromBlock, 10))
	}

	if filters.ToBlock > 0 {
		q.Set("to_block", strconv.FormatInt(filters.ToBlock, 10))
	}

	u.RawQuery = q.Encode()

	return u, nil
}

func (c *Client) Get(url string) (*Response, error) {
	return c.Do(http.MethodGet, url, http.NoBody)
}

func (c *Client) Post(url string, body io.Reader) (*Response, error) {
	return c.Do(http.MethodPost, url, body)
}

func (c *Client) Do(method string, url string, body io.Reader) (*Response, error) {
	req, _ := http.NewRequest(method, url, body)

	req.Header.Set("Content-Type", "application/json;")
	req.Header.Set("X-API-KEY", c.APIKey)

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
