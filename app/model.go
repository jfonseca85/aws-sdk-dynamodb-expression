package app

type Model struct {
	ID       string `json:"id" dynamodbav:"ID"`
	Version  string `json:"version,omitempty" dynamodbav:"Version"`
	Document string `json:"document,omitempty" dynamodbav:"document"`
	ArnAsf   string `json:"arn-asf" dynamodbav:"arnAsf"`
	Status   string `json:"status" dynamodbav:"status"`
}
