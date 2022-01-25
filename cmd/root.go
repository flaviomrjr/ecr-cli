/*
Copyright © 2022 Flavio Rocha

*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// this func gets repo name
func SetRepo(cmd *cobra.Command, args []string) string {
	// get repo name
	repo, err := cmd.Flags().GetString("repo")
	if err != nil {
		log.Fatal(err)
	}

	return repo
}

// this func gets profile name or set profile to shared
func SetProfile(cmd *cobra.Command, args []string) {
	// set profile
	profile, err := cmd.Flags().GetString("profile")
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("AWS_PROFILE", profile)
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
	rootCmd.PersistentFlags().StringP("repo", "r", "", "Set repo name")
	rootCmd.PersistentFlags().StringP("profile", "p", "shared", "Set AWS profile")
}
