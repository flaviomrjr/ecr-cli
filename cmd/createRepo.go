/*
Copyright © 2022 Flavio Rocha flavio.rocha16@gmail.com

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

// convert bytes to string
func bytesToString(data []byte) string {
	return string(data[:])
}

// get json policy
func GetPolicy(cmd *cobra.Command, args []string) string {
	// get json path
	jsonPath, err := cmd.Flags().GetString("policy-file")

	// read json file
	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	// convert var jsonData from byte to string
	policy := bytesToString(jsonData)

	return policy
}

// set ecr policy
func SetPolicy(policy string, repo string) {

	resultPolicy, err := NewSession().SetRepositoryPolicy(&ecr.SetRepositoryPolicyInput{
		PolicyText:     &policy,
		RepositoryName: &repo,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultPolicy)
}

// create ecr repo
func CreateRepoECR(repo string) {

	result, err := NewSession().CreateRepository(&ecr.CreateRepositoryInput{
		RepositoryName:             aws.String(repo),
		ImageScanningConfiguration: &ecr.ImageScanningConfiguration{ScanOnPush: aws.Bool(true)},
		ImageTagMutability:         aws.String(ecr.ImageTagMutabilityMutable),
		EncryptionConfiguration:    &ecr.EncryptionConfiguration{EncryptionType: aws.String("AES256")},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

// createRepoCmd represents the create-repo command
var createRepoCmd = &cobra.Command{
	Use:   "create-repo",
	Short: "Add a new repo in ECR",
	Long:  `This function creates one or more ECR repos`,
	Run: func(cmd *cobra.Command, args []string) {
		// set aws region
		SetRegion(cmd, args)

		// set aws profile
		SetProfile(cmd, args)

		// set repo
		repos := SetRepo(cmd, args)

		if noPolicy, err := cmd.Flags().GetBool("no-policy"); err != nil {
			log.Fatal(err)
		} else {
			if noPolicy == false {
				// get repo policy
				policy := GetPolicy(cmd, args)

				for _, repo := range repos {
					// create ecr repo
					CreateRepoECR(repo)

					// set repo policy
					SetPolicy(policy, repo)
				}
			} else {
				for _, repo := range repos {
					// create repo without policy
					CreateRepoECR(repo)
				}
			}
		}
	},
}

func init() {
	createRepoCmd.Flags().StringP("policy-file", "f", "policy.json", "Set Policy Path. e.g.: ~/Documents/policy.json")
	createRepoCmd.Flags().BoolP("no-policy", "", false, "Do not apply repo policy")
	rootCmd.AddCommand(createRepoCmd)
}
