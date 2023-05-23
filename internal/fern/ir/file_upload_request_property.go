package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type FileUploadRequestProperty struct {
	Type         string
	File         *FileProperty
	BodyProperty *InlinedRequestBodyProperty
}

func (x *FileUploadRequestProperty) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	x.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "file":
		value := new(FileProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.File = value
	case "bodyProperty":
		value := new(InlinedRequestBodyProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.BodyProperty = value
	}
	return nil
}

type FileUploadRequestPropertyVisitor interface {
	VisitFile(*FileProperty) error
	VisitBodyProperty(*InlinedRequestBodyProperty) error
}

func (x *FileUploadRequestProperty) Accept(v FileUploadRequestPropertyVisitor) error {
	switch x.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "file":
		return v.VisitFile(x.File)
	case "bodyProperty":
		return v.VisitBodyProperty(x.BodyProperty)
	}
}