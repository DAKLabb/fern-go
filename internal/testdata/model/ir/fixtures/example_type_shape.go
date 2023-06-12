// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ExampleTypeShape struct {
	Type   string
	Alias  *ExampleAliasType
	Enum   *ExampleEnumType
	Object *ExampleObjectType
	Union  *ExampleSingleUnionType
}

func (e *ExampleTypeShape) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "alias":
		value := new(ExampleAliasType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Alias = value
	case "enum":
		value := new(ExampleEnumType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Enum = value
	case "object":
		value := new(ExampleObjectType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Object = value
	case "union":
		value := new(ExampleSingleUnionType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Union = value
	}
	return nil
}

func (e ExampleTypeShape) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "alias":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleAliasType
		}{
			Type:             e.Type,
			ExampleAliasType: e.Alias,
		}
		return json.Marshal(marshaler)
	case "enum":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleEnumType
		}{
			Type:            e.Type,
			ExampleEnumType: e.Enum,
		}
		return json.Marshal(marshaler)
	case "object":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleObjectType
		}{
			Type:              e.Type,
			ExampleObjectType: e.Object,
		}
		return json.Marshal(marshaler)
	case "union":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleSingleUnionType
		}{
			Type:                   e.Type,
			ExampleSingleUnionType: e.Union,
		}
		return json.Marshal(marshaler)
	}
}

type ExampleTypeShapeVisitor interface {
	VisitAlias(*ExampleAliasType) error
	VisitEnum(*ExampleEnumType) error
	VisitObject(*ExampleObjectType) error
	VisitUnion(*ExampleSingleUnionType) error
}

func (e *ExampleTypeShape) Accept(v ExampleTypeShapeVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "alias":
		return v.VisitAlias(e.Alias)
	case "enum":
		return v.VisitEnum(e.Enum)
	case "object":
		return v.VisitObject(e.Object)
	case "union":
		return v.VisitUnion(e.Union)
	}
}