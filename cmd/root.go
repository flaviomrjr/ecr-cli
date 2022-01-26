/*
Copyright © 2022 Flavio Rocha

*/
package cmd

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

// create a new session
func NewSession() *ecr.ECR {
	svc := ecr.New(session.New())

	return svc
}

// set repo name
func SetRepo(cmd *cobra.Command, args []string) []string {

	repo, err := cmd.Flags().GetStringSlice("repo")
	if err != nil {
		log.Fatal(err)
	}

	return repo
}

// this func gets profile name or set profile to shared
func SetProfile(cmd *cobra.Command, args []string) {

	profile, err := cmd.Flags().GetString("profile")
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("AWS_PROFILE", profile)
}

// set aws region
func SetRegion(cmd *cobra.Command, args []string) {

	region, err := cmd.Flags().GetString("region")
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("AWS_REGION", region)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ecr",
	Short: "Application tool used to manage ECR docker repos",
	Long:  `This cli was built to help users manage their ECR repos. `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// global flags
	rootCmd.PersistentFlags().StringSliceP("repo", "r", []string{}, "Set repo name")
	rootCmd.PersistentFlags().StringP("profile", "p", "shared", "Set AWS profile")
	rootCmd.PersistentFlags().StringP("region", "", "sa-east-1", "Set AWS region")
}
