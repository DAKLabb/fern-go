package ir

type ExampleSingleUnionType struct {
	WireDiscriminantValue string                            `json:"wireDiscriminantValue"`
	Properties            *ExampleSingleUnionTypeProperties `json:"properties"`
}