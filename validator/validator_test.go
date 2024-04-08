package validator

import (
	"testing"
)

type ValidatorTest struct {
	JSON     []byte
	expected interface{}
}

var resourceValidateTests = []ValidatorTest{
	{[]byte(`{"PolicyName": "test","PolicyDocument": {"Statement": []}}`), true},
	{[]byte(`{"PolicyName": "test", "PolicyDocument": {"Statement": [{"Resource": "arn:aws:s3:::*"}]}}`), true},
	{[]byte(`{"PolicyName": "test", "PolicyDocument": {"Statement": [{"Resource": "*"}]}}`), false},
	{[]byte(`{"PolicyName": "test", "PolicyDocument": {"Statement": [{"Resource": ["arn:aws:s3:::test"]}]}}`), true},
	{[]byte(`{"PolicyName": "test", "PolicyDocument": {"Statement": [{"Resource": "*"}, {"Resource": "arn:aws:s3:::test"}]}}`), false},
	{[]byte(`{"PolicyName": "test", "PolicyDocument": {"Statement": [{"Resource": ["*", "arn:aws:s3:::test"]}, {"Resource": "arn:aws:s3:::test"}]}}`), false},
}

func TestResourceValidate(t *testing.T) {
	for _, test := range resourceValidateTests {
		policy := &AWSRolePolicy{}

		isValid := policy.VerifyAWSRolePolicy(test.JSON)

		if isValid != test.expected {
			t.Errorf("Expected %t, but got %t", test.expected, isValid)
		}
	}
}
