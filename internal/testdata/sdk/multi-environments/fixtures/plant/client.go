// This file was auto-generated by Fern from our API Definition.

package plant

import (
	context "context"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/multi-environments/fixtures/core"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

func (c *Client) GetPlant(ctx context.Context) (string, error) {
	baseURL := "https://plants.yoursite.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "plants"

	var response string
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
