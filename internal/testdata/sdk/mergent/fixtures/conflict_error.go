// Generated by Fern. Do not edit.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type ConflictError struct {
	StatusCode int
	Body       *Error
}

func (c *ConflictError) Error() string {
	return fmt.Sprintf("409: %+v", *c)
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	body := new(Error)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}