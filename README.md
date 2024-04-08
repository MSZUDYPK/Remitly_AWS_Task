# Remitly AWS Task
## Requirements
- [Go (1.22)](https://golang.org/dl/)

## Installation
Clone the repository:
```bash
git clone https://github.com/MSZUDYPK/Remitly_AWS_Task.git
``` 
Change directory to the project:
```bash
cd Remitly_AWS_Task
```
To compile and build the application, use the following command:
```bash
go build
```
## Usage
> [!IMPORTANT]  
> The expected input format is defined as valid [AWS::IAM::Role Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html).
> 
> See more at: [AWS IAM Policy Validator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-policy-validation.html)
### Application
The application is a simple command line tool that validates an AWS IAM Role Policy. 
The application reads a JSON file containing the IAM Role Policy and checks statements Resource field(s). 
The application will return `false` if the resource field contain single asterisk and `true` otherwise.

To run the application, use the following command:
```bash
./remitly_aws_task <path_to_json_file>
```

### Package
Here is a basic example of how to use the `VerifyAWSRolePolicy`:

```go
package main

import (
	"fmt"
	"remitly_aws_task/validator"
)

func main() {
	var buffer = []byte(`{"PolicyName": "test", "PolicyDocument": {"Statement": [{"Resource": "*"}]}}`)

	var policy = validator.AWSRolePolicy{}
	isValid := policy.VerifyAWSRolePolicy(buffer)

	fmt.Println(isValid)
}
```

## Running Unit Tests
To run the unit tests, use the following command:
```bash
go test ./...
```