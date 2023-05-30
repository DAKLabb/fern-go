package ir

type UnionTypeDeclaration struct {
	Discriminant *NameAndWireValue `json:"discriminant"`
	// A list of other types to inherit from
	Extends        []*DeclaredTypeName `json:"extends"`
	Types          []*SingleUnionType  `json:"types"`
	BaseProperties []*ObjectProperty   `json:"baseProperties"`
}
