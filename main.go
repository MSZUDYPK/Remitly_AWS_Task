package main

import (
	"fmt"
	"os"
	"remitly_aws_task/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file as an argument.")
		return
	}

	buffer, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	policy := &validator.AWSRolePolicy{}
	isValid := policy.VerifyAWSRolePolicy(buffer)

	fmt.Println(isValid)
}
