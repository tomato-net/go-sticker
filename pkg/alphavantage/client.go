package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://www.alphavantage.co/query"
	defaultTimeout = 5 * time.Second
)

type Client struct {
	apiKey     string
	baseURL    string
	retries    int
	httpClient *http.Client
}

type ClientOptions struct {
	APIKey     string
	BaseURL    string
	Retries    int
	HTTPClient *http.Client
}

func NewClient(options ClientOptions) (*Client, error) {
	if strings.TrimSpace(options.APIKey) == "" {
		return nil, fmt.Errorf("no API key provided")
	}

	client := &Client{
		apiKey:  options.APIKey,
		baseURL: defaultBaseURL,
		retries: options.Retries,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	if options.BaseURL != "" {
		client.baseURL = options.BaseURL
	}

	if options.HTTPClient != nil {
		client.httpClient = options.HTTPClient
	}

	return client, nil
}

func (c *Client) Do(ctx context.Context, params url.Values, into interface{}) error {
	var lastError error
	for retries := 0; retries <= c.retries; retries++ {
		rawResp, err := c.do(ctx, params)
		if err != nil {
			// TODO: Handle rate limiting with backoff
			lastError = err
			continue
		}

		body, err := io.ReadAll(rawResp.Body)
		if err != nil {
			return fmt.Errorf("io.ReadAll: %w", err)
		}

		errorCheck := ErrorMessage{}
		json.Unmarshal(body, &errorCheck)
		if IsError(errorCheck) {
			return errorCheck.Error()
		}

		if err := json.Unmarshal(body, into); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}

		return nil
	}

	return fmt.Errorf("failed after %d retries, last error: %w", c.retries, lastError)
}

func (c *Client) do(ctx context.Context, params url.Values) (*http.Response, error) {
	params.Set(ParamKeyAPIKey, c.apiKey)
	params.Set(ParamKeyDataType, DataTypeJSON)
	url := fmt.Sprintf("%s?%s", c.baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("User-Agent", "Go Stock Ticker/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do: %w", err)
	}

	return resp, nil
}
