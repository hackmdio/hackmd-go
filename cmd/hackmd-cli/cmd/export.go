/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteId, _ := cmd.Flags().GetString("noteId")

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

	exportCmd.PersistentFlags().String("noteId", "", "The ID of the note to export")
}
