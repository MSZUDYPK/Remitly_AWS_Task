package main

import (
	"os"
	"remitly_aws_task/validator"
	"testing"
)

type ValidatorTest struct {
	filePath string
	expected interface{}
}

var resourceValidateTests = []ValidatorTest{
	{"./resources/noResourceFieldJSON", true},
	{"./resources/SingleAsteriskResourceFieldJSON", false},
	{"./resources/SingleAsteriskMultipleResourceFieldsJSON", false},
	{"./resources/MultipleAsteriskResourceFieldJSON", true},
	{"./resources/noAsteriskResourceFieldJSON", true},
}

func TestResourceValidate(t *testing.T) {
	for _, test := range resourceValidateTests {
		buffer, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Errorf("Error reading file: %s", err)
		}

		policy := &validator.AWSRolePolicy{}

		isValid := policy.VerifyAWSRolePolicy(buffer)

		if isValid != test.expected {
			t.Errorf("Expected %t, but got %t", test.expected, isValid)
		}
	}
}
