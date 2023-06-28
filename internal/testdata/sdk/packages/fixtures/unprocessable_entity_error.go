// Generated by Fern. Do not edit.

package api

import (
	json "encoding/json"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
)

type UnprocessableEntityError struct {
	*core.APIError
	Body *Error
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	body := new(Error)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}