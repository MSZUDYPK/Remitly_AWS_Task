package validator

import (
	"encoding/json"
	"log"
)

type AWSRolePolicy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Resource interface{} `json:"Resource"`
}

func (statement *Statement) hasSingleAsteriskInResource() bool {
	switch res := statement.Resource.(type) {
	case string:
		if res == "*" {
			return true
		}
	case []interface{}:
		for _, r := range res {
			if str, ok := r.(string); ok && str == "*" {
				return true
			}
		}
	}
	return false
}

func (policyDocument *PolicyDocument) hasValidStatementsResourceFields() bool {
	for _, statement := range policyDocument.Statement {
		if statement.hasSingleAsteriskInResource() {
			return false
		}
	}
	return true
}

func (policy *AWSRolePolicy) VerifyAWSRolePolicy(buffer []byte) bool {
	err := json.Unmarshal(buffer, policy)
	if err != nil {
		log.Fatal(err)
	}

	return policy.PolicyDocument.hasValidStatementsResourceFields()
}
