package api

import (
	uuid "github.com/gofrs/uuid"
)

// This is a Foo.
type Foo struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
