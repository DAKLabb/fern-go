// This file was auto-generated by Fern from our API Definition.

package core

import (
	base64 "encoding/base64"
	fmt "fmt"
	uuid "github.com/google/uuid"
	http "net/http"
	time "time"
)

// ClientOption adapts the behavior of the generated client.
type ClientOption func(*ClientOptions)

// ClientOptions defines all of the possible client options.
// This type is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
type ClientOptions struct {
	BaseURL              string
	HTTPClient           HTTPClient
	HTTPHeader           http.Header
	Custom               *[]byte
	XApiName             string
	XApiId               uuid.UUID
	XApiDatetime         time.Time
	XApiDate             time.Time
	XApiBytes            []byte
	XApiOptionalName     *string
	XApiOptionalId       *uuid.UUID
	XApiOptionalDatetime *time.Time
	XApiOptionalDate     *time.Time
	XApiOptionalBytes    *[]byte
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
func (c *ClientOptions) ToHeader() http.Header {
	header := c.cloneHeader()
	if c.Custom != nil {
		header.Set("X-API-Custom-Key", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(*c.Custom)))
	}
	header.Set("X-API-Name", fmt.Sprintf("%v", c.XApiName))
	header.Set("X-API-ID", fmt.Sprintf("%v", c.XApiId))
	header.Set("X-API-Datetime", fmt.Sprintf("%v", c.XApiDatetime.Format(time.RFC3339)))
	header.Set("X-API-Date", fmt.Sprintf("%v", c.XApiDate.Format("2006-01-02")))
	header.Set("X-API-Bytes", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(c.XApiBytes)))
	if c.XApiOptionalName != nil {
		header.Set("X-API-Optional-Name", fmt.Sprintf("%v", *c.XApiOptionalName))
	}
	if c.XApiOptionalId != nil {
		header.Set("X-API-Optional-ID", fmt.Sprintf("%v", *c.XApiOptionalId))
	}
	if c.XApiOptionalDatetime != nil {
		header.Set("X-API-Optional-Datetime", fmt.Sprintf("%v", c.XApiOptionalDatetime.Format(time.RFC3339)))
	}
	if c.XApiOptionalDate != nil {
		header.Set("X-API-Optional-Date", fmt.Sprintf("%v", c.XApiOptionalDate.Format("2006-01-02")))
	}
	if c.XApiOptionalBytes != nil {
		header.Set("X-API-Optional-Bytes", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(*c.XApiOptionalBytes)))
	}
	header.Set("X-API-Fern-Header", fmt.Sprintf("%v", "fern"))
	return header
}

func (c *ClientOptions) cloneHeader() http.Header {
	return c.HTTPHeader.Clone()
}
