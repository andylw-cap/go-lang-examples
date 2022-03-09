package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// Usage:
// go run sts_assume_role.go
func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	if err != nil {
		fmt.Println("NewSession Error", err)
		return
	}
	// Ask for 2FA token code required to auth to role
	var tokenCode string
	fmt.Print("AWS 2FA Token: ")
	fmt.Scanf("%s", &tokenCode)

	// Create a STS client
	svc := sts.New(sess)

	roleToAssumeArn := "arn:aws:iam::904806826062:role/OrganizationEngineerAccessRole"
	sessionName := "terraform"
	mfaSerialArn := "arn:aws:iam::843361875856:mfa/andrewwaters"
	result, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         &roleToAssumeArn,
		RoleSessionName: &sessionName,
		SerialNumber:    &mfaSerialArn,
		TokenCode:       &tokenCode,
	})

	if err != nil {
		fmt.Println("AssumeRole Error", err)
		return
	}

	fmt.Println(result.AssumedRoleUser)
	fmt.Println(result.Credentials)

	awsKey := result.Credentials.AccessKeyId
	awsSecret := result.Credentials.SecretAccessKey
	awsToken := result.Credentials.SessionToken

	os.Setenv("AWS_ACCESS_KEY_ID", *awsKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", *awsSecret)
	os.Setenv("AWS_SESSION_TOKEN", *awsToken)

	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

}
