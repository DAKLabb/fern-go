// This file was auto-generated by Fern from our API Definition.

package api

type AnotherType struct {
	String *string `json:"string,omitempty"`
	Type   *Type   `json:"type,omitempty"`
}

type Type struct {
	Name string `json:"name"`
}
