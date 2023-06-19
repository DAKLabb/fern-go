// Generated by Fern. Do not edit.

package api

import (
	context "context"
	errors "errors"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params/fixtures/core"
	io "io"
	http "net/http"
)

type getAllUsersEndpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newGetAllUsersEndpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *getAllUsersEndpoint {
	return &getAllUsersEndpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (g *getAllUsersEndpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (g *getAllUsersEndpoint) Call(ctx context.Context, request *GetAllUsersRequest) (string, error) {
	endpointURL := g.url
	var response string
	if err := core.DoRequest(
		ctx,
		g.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		response,
		g.header,
		g.decodeError,
	); err != nil {
		return response, err
	}
	return response, nil
}