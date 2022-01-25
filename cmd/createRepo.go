/*
Copyright © 2022 Flavio Rocha flavio.rocha16@gmail.com

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

// convert bytes to string
func BytesToString(data []byte) string {
	return string(data[:])
}

// configure repo policy
func SetPolicy(cmd *cobra.Command, args []string) string {
	// read json file
	jsonPath, err := cmd.Flags().GetString("policy-file")

	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	// convert var jsonData from byte to string
	policy := BytesToString(jsonData)

	return policy
}

// create a ECR repo
func CreateRepoECR(cmd *cobra.Command, args []string, repo string) {
	// get repo policy
	policy := SetPolicy(cmd, args)

	// create a new session
	svc := ecr.New(session.New())

	// create repo and show the result
	result, err := svc.CreateRepository(&ecr.CreateRepositoryInput{
		RepositoryName:             aws.String(repo),
		ImageScanningConfiguration: &ecr.ImageScanningConfiguration{ScanOnPush: aws.Bool(true)},
		ImageTagMutability:         aws.String(ecr.ImageTagMutabilityMutable),
		EncryptionConfiguration:    &ecr.EncryptionConfiguration{EncryptionType: aws.String("AES256")},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	// create policy and show the result
	resultPolicy, err := svc.SetRepositoryPolicy(&ecr.SetRepositoryPolicyInput{
		PolicyText:     &policy,
		RepositoryName: &repo,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultPolicy)
}

// createRepoCmd represents the create-repo command
var createRepoCmd = &cobra.Command{
	Use:   "create-repo",
	Short: "Add a new repo in ECR",
	Long:  `This function creates a new image repo in ECR with some configs imported from a json-file`,
	Run: func(cmd *cobra.Command, args []string) {
		// call func SetProfile
		SetProfile(cmd, args)

		// call func CreateRepoECR
		CreateRepoECR(cmd, args, SetRepo(cmd, args))
	},
}

func init() {
	createRepoCmd.Flags().StringP("policy-file", "f", "policy.json", "Set Policy Path. e.g.: ~/Documents/policy.json")
	rootCmd.AddCommand(createRepoCmd)
}
