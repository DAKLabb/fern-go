package api

type Bar struct {
	// This is a Foo field.
	Foo *Foo `json:"foo"`
}
