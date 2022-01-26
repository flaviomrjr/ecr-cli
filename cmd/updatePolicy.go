/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updatePolicyCmd represents the updatePolicy command
var updatePolicyCmd = &cobra.Command{
	Use:   "update-policy",
	Short: "Update repo policy",
	Long:  `This function updates the repo policy`,
	Run: func(cmd *cobra.Command, args []string) {
		// set aws region
		SetRegion(cmd, args)

		// set aws profile
		SetProfile(cmd, args)

		// set repos
		repos := SetRepo(cmd, args)

		// get repo policy
		policy := GetPolicy(cmd, args)

		for _, repo := range repos {
			// update policy
			SetPolicy(policy, repo)
		}
	},
}

func init() {
	updatePolicyCmd.Flags().StringP("policy-file", "f", "", "Set Policy Path. e.g.: ~/Documents/policy.json")
	updatePolicyCmd.MarkFlagRequired("policy-file")
	rootCmd.AddCommand(updatePolicyCmd)
}
