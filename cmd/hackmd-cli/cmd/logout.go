/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from your HackMD account",
	Long: `The 'logout' command is used to log out from your HackMD account on the 
command line interface (CLI). This command effectively removes your locally stored 
access token, which means you'll have to use the 'login' command and enter your 
access token again to authenticate the next time you want to perform any operation 
that requires authentication.

After running this command, your HackMD access token will be removed from the 
locally stored configuration and a message confirming your logout will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.RemoveAccessTokenFromConfig()

		fmt.Println("You have been logged out.")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
