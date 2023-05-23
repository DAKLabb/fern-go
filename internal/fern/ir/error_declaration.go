package ir

type ErrorDeclaration struct {
	Docs              *string            `json:"docs"`
	Name              *DeclaredErrorName `json:"name"`
	DiscriminantValue *NameAndWireValue  `json:"discriminantValue"`
	Type              *TypeReference     `json:"type"`
	StatusCode        int                `json:"statusCode"`
}