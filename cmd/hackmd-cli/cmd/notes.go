/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		notes, err := api.GetNoteList()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, note := range *notes {
			fmt.Println(note.ID, note.Title)
		}
	},
}

var notesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note",
	Long: `Create a new note. The note will be created with the title provided. If no title is provided, a random title will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		title, _ := cmd.Flags().GetString("title")
		content, _ := cmd.Flags().GetString("content")

		// TODO: add permission fields and validation

		note, err := api.CreateNote(&HackMDClient.CreateNoteOptions{
			Title: title,
			Content: content,
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(note.ID, note.Title)
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)

	notesCmd.AddCommand(notesCreateCmd)
	notesCreateCmd.PersistentFlags().String("title", "", "The title of the note to create")
	notesCreateCmd.PersistentFlags().String("content", "", "The content of the note to create")
}
