package verifier

import (
	"encoding/json"
	"log/slog"

	"github.com/migonov/json_verifier/utils"
)

type (
	PolicyDocumentStatement struct {
		Resource *utils.StringOrSlice `json:"Resource,omitempty"`
	}

	PolicyDocument struct {
		Statements []PolicyDocumentStatement `json:"Statement"`
	}

	AWSIAMRolePolicy struct {
		PolicyDocument PolicyDocument
	}
)

func Verify(inputJSON []byte) bool {
	var policy AWSIAMRolePolicy
	if err := json.Unmarshal(inputJSON, &policy); err != nil {
		slog.Error("Error parsing input JSON AWS:IAM:Role Policy document")
		panic(err)
	}

	statements := policy.PolicyDocument.Statements

	for _, statement := range statements {
		if statement.Resource != nil && containsAsterisk(statement.Resource.Values) {
			return false
		}
	}

	return true
}

func containsAsterisk(resources []string) bool {
	for _, resource := range resources {
		if resource == "*" {
			slog.Info("Resource contains single asterisk")
			return true
		}
	}
	return false
}
