package ir

type ObjectTypeDeclaration struct {
	// A list of other types to inherit from
	Extends    []*DeclaredTypeName `json:"extends"`
	Properties []*ObjectProperty   `json:"properties"`
}
