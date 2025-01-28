/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createRepositoryCmd represents the createRepository command
var createRepositoryCmd = &cobra.Command{
	Use:   "create-repository",
	Short: "This process create a repository remote and local.",
	Long: `This is responsible to create a repository remote in Github, Bitbucket and more;
	and in your local machine too.Doing the bind among them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createRepository called")
	},
}

func init() {
	rootCmd.AddCommand(createRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRepositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRepositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
