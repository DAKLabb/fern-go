package core

import (
	http "net/http"
)

// ---
// This file is not used as a template for code generation. It is only included so that
// the client_test.go template can actually run as a real test file in this repository.
// ---

// ClientOption adapts the behavior of the generated client.
type ClientOption func(*ClientOptions)

// ClientOptions defines all of the possible client options.
// This type is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
type ClientOptions struct {
	BaseURL    string
	HTTPClient HTTPClient
	HTTPHeader http.Header
}

// NewClientOptions returns a new *ClientOptions value.
// This function is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		HTTPClient: http.DefaultClient,
		HTTPHeader: make(http.Header),
	}
}

// ToHeader maps the configured client options into a http.Header issued
// on every request.
func (c *ClientOptions) ToHeader() http.Header { return c.cloneHeader() }

func (c *ClientOptions) cloneHeader() http.Header {
	return c.HTTPHeader.Clone()
}
