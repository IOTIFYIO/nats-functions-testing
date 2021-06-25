package main

// Function defines model for Function.
type Function struct {
	FunctionDefinition string `json:"functionDefinition"`
	Name               string `json:"name"`
	NatsSubjectTrigger string `json:"natsSubjectTrigger"`
	Id                 string `json:"id"`
}
