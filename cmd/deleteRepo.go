/*
Copyright © 2022 Flavio Rocha flavio.rocha16@gmail.com

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

// deleteRepoCmd represents the deleteRepo command
var deleteRepoCmd = &cobra.Command{
	Use:   "delete-repo",
	Short: "Delete ECR repo",
	Long:  `This function deletes one or more ECR repos`,
	Run: func(cmd *cobra.Command, args []string) {
		// set aws region
		SetRegion(cmd, args)

		// set aws profile
		SetProfile(cmd, args)

		// set repos
		repos := SetRepo(cmd, args)

		for _, repo := range repos {
			result, err := NewSession().DeleteRepository(&ecr.DeleteRepositoryInput{
				Force:          aws.Bool(true),
				RepositoryName: &repo,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
			fmt.Println("Repo deleted")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteRepoCmd)
}
