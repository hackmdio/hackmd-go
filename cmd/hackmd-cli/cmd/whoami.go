/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/spf13/cobra"
)

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display logged-in user's information",
	Long:  `The 'whoami' command displays the information of the currently logged-in user in the HackMD account. If the user is not logged in, the command will return a message prompting the user to log in.`,
	Run: func(cmd *cobra.Command, args []string) {
		if internal.IsAccessTokenPresent() {
			internal.GetMeFlow(false)
		} else {
			fmt.Println("You are not logged in. Please run `hackmd-cli login` to log in first.")
		}
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
