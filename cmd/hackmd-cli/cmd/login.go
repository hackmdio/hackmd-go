/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your HackMD account",
	Long: `The 'login' command is used to authenticate and gain access to your HackMD account 
from the command line interface (CLI). This command requires you to enter an access token 
for authentication. 

To use the 'login' command:

  $ hackmd-cli login

You will be prompted to enter your HackMD access token. After successful authentication, 
you will be logged in to your HackMD account and your account information will be displayed.

It's important to keep your access token secure. Avoid using the 'login' command in scripts 
or automated processes where your token might be exposed, and never store your token in 
plain text.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetMeFlow(true)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
