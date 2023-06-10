/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/cmd/hackmd-cli/internal"
	"github.com/hackmdio/hackmd-go/cmd/hackmd-cli/internal/flags"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Retrieve and display the user's HackMD note history",

	Long: `The 'history' command is used to retrieve and display the list of HackMD notes from the user's history. It provides an overview of all the notes that the user has interacted with in the past.

This command fetches and displays the notes' history in the user's console. Optionally, you can specify an output format to structure the display of these notes. The command handles potential errors gracefully, ensuring a user-friendly experience.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		notes, err := api.GetHistory()
		if err != nil {
			fmt.Println(err)
			return
		}

		output := cmd.Flag("output").Value.String()
		internal.PrintNotes(output, notes)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	flags.AddCommandFlags(historyCmd, []flags.FlagData{flags.OutputFlag})
}
