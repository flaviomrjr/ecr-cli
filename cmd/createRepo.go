/*
Copyright © 2022 Flavio Rocha flavio.rocha16@gmail.com

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

// set json policy
var jsonData = `{
	"Version": "2008-10-17",
	"Statement": [
	  {
		"Sid": "eks_access",
		"Effect": "Allow",
		"Principal": {
		  "AWS": [
			"arn:aws:iam::000000000000:root",
			"arn:aws:iam::111111111111:root"
		  ]
		},
		"Action": [
		  "ecr:BatchCheckLayerAvailability",
		  "ecr:BatchGetImage",
		  "ecr:CompleteLayerUpload",
		  "ecr:DescribeImages",
		  "ecr:GetDownloadUrlForLayer",
		  "ecr:GetLifecyclePolicy",
		  "ecr:GetLifecyclePolicyPreview",
		  "ecr:GetRepositoryPolicy",
		  "ecr:ListImages",
		  "ecr:DescribeRepositories",
		  "ecr:InitiateLayerUpload",
		  "ecr:UploadLayerPart",
		  "ecr:PutImage",
		  "ecr:PutLifecyclePolicy",
		  "ecr:SetRepositoryPolicy",
		  "ecr:StartLifecyclePolicyPreview"
		]
	  }
	]
  }`

// createRepoCmd represents the create-repo command
var createRepoCmd = &cobra.Command{
	Use:   "create-repo",
	Short: "Add a new repo in ECR",
	Long:  `This function creates a new image repo in ECR with some configs imported from a json-file`,
	Run: func(cmd *cobra.Command, args []string) {
		// set region
		os.Setenv("AWS_REGION", "sa-east-1")

		// read aws credentials file
		os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

		// get repo name
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			log.Fatal(err)
		}

		// set profile
		profile, err := cmd.Flags().GetString("profile")
		if err != nil {
			log.Fatal(err)
		}
		os.Setenv("AWS_PROFILE", profile)

		// create a new session
		svc := ecr.New(session.New())
		if err != nil {
			log.Fatal(err)
		}

		// set repo parameters
		input := &ecr.CreateRepositoryInput{
			RepositoryName:             aws.String(repo),
			ImageScanningConfiguration: &ecr.ImageScanningConfiguration{ScanOnPush: aws.Bool(true)},
			ImageTagMutability:         aws.String(ecr.ImageTagMutabilityMutable),
			EncryptionConfiguration:    &ecr.EncryptionConfiguration{EncryptionType: aws.String("AES256")},
		}

		// create repo and show the result
		result, err := svc.CreateRepository(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ecr.ErrCodeServerException:
					fmt.Println(ecr.ErrCodeServerException, aerr.Error())
				case ecr.ErrCodeInvalidParameterException:
					fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
				case ecr.ErrCodeInvalidTagParameterException:
					fmt.Println(ecr.ErrCodeInvalidTagParameterException, aerr.Error())
				case ecr.ErrCodeTooManyTagsException:
					fmt.Println(ecr.ErrCodeTooManyTagsException, aerr.Error())
				case ecr.ErrCodeRepositoryAlreadyExistsException:
					fmt.Println(ecr.ErrCodeRepositoryAlreadyExistsException, aerr.Error())
				case ecr.ErrCodeLimitExceededException:
					fmt.Println(ecr.ErrCodeLimitExceededException, aerr.Error())
				case ecr.ErrCodeKmsException:
					fmt.Println(ecr.ErrCodeKmsException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)

		// set policy parameters
		imputPolicy := &ecr.SetRepositoryPolicyInput{
			PolicyText:     &jsonData,
			RepositoryName: &repo,
		}

		// create policy and show the result
		resultPolicy, err := svc.SetRepositoryPolicy(imputPolicy)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resultPolicy)
	},
}

func init() {
	createRepoCmd.Flags().StringP("repo", "r", "", "Create a new image repo in ECR")
	createRepoCmd.Flags().StringP("profile", "p", "shared", "Set AWS profile")
	rootCmd.AddCommand(createRepoCmd)
}
