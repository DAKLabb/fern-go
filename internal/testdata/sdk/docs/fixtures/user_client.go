// This file was auto-generated by Fern from our API Definition.

package api

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/docs/fixtures/core"
	http "net/http"
	url "net/url"
)

type UserClient interface {
	GetName(ctx context.Context, userId string, request *GetNameRequest) (string, error)
}

func NewUserClient(opts ...core.ClientOption) UserClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &userClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type userClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Returns the username associated with the given userId.
//
// userId uniquely identifies a user.
func (u *userClient) GetName(ctx context.Context, userId string, request *GetNameRequest) (string, error) {
	baseURL := ""
	if u.baseURL != "" {
		baseURL = u.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"/users/%v/get-name", userId)

	queryParams := make(url.Values)
	var filterDefaultValue string
	if request.Filter != filterDefaultValue {
		queryParams.Add("filter", fmt.Sprintf("%v", request.Filter))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := u.header.Clone()
	var xEndpointHeaderDefaultValue string
	if request.XEndpointHeader != xEndpointHeaderDefaultValue {
		headers.Add("X-Endpoint-Header", fmt.Sprintf("%v", request.XEndpointHeader))
	}

	var response string
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		headers,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}