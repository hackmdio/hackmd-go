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

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a specific HackMD note's content",
	Long: `The 'export' command retrieves and displays the content of a specific HackMD note, identified by its unique ID. This can be useful when you want to manipulate or save the note's content outside of the HackMD environment.

To execute this command, you must provide the ID of the note you wish to export. If the provided ID corresponds to an existing note, the content of the note will be displayed. 

Note that you must be logged in and have the appropriate permissions to access the note you're trying to export.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteId, _ := cmd.Flags().GetString(flags.NoteIDFlag.Name)

		if noteId == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		note, err := api.GetNote(noteId)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(note.Content)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	flags.AddCommandFlags(exportCmd, []flags.FlagData{flags.NoteIDFlag})
}
