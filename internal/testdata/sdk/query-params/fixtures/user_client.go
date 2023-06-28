// Generated by Fern. Do not edit.

package api

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params/fixtures/core"
	http "net/http"
	url "net/url"
	strings "strings"
)

type UserClient interface {
	GetAllUsers(ctx context.Context, request *GetAllUsersRequest) (string, error)
}

func NewUserClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) UserClient {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	return &userClient{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		header:     options.ToHeader(),
	}
}

type userClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (u *userClient) GetAllUsers(ctx context.Context, request *GetAllUsersRequest) (string, error) {
	headers := u.header.Clone()
	var xEndpointHeaderDefaultValue string
	if request.XEndpointHeader != xEndpointHeaderDefaultValue {
		headers.Add("X-Endpoint-Header", fmt.Sprintf("%v", request.XEndpointHeader))
	}

	endpointURL := u.baseURL + "/" + "/users/all"
	queryParams := make(url.Values)
	var limitDefaultValue *int
	if request.Limit != limitDefaultValue {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
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