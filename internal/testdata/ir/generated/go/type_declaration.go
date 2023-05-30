package ir

// A type, which is a name and a shape
type TypeDeclaration struct {
	Docs         *string           `json:"docs"`
	Availability *Availability     `json:"availability"`
	Name         *DeclaredTypeName `json:"name"`
	Shape        *Type             `json:"shape"`
	Examples     []*ExampleType    `json:"examples"`
	// All other named types that this type references (directly or indirectly)
	ReferencedTypes []*DeclaredTypeName `json:"referencedTypes"`
}
