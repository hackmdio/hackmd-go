/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/hackmdio/hackmd-go/hackmd-cli/internal/flags"
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

		output := cmd.Flag("output").Value.String()
		internal.PrintNotes(output, notes)
	},
}

func processInput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	var pipeText string
	for scanner.Scan() {
		pipeText += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return pipeText
}

var notesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note",
	Long:  `Create a new note. The note will be created with the title provided. If no title is provided, a random title will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		title, _ := cmd.Flags().GetString("title")
		content, _ := cmd.Flags().GetString("content")
		readPermission, _ := cmd.Flags().GetString("readPermission")
		writePermission, _ := cmd.Flags().GetString("writePermission")
		commentPermission, _ := cmd.Flags().GetString("commentPermission")

		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeNamedPipe) != 0 {
			// if no content
			if content == "" {
				content = processInput(os.Stdin)
			}
		}

		note, err := api.CreateNote(&HackMDClient.CreateNoteOptions{
			Title:             title,
			Content:           content,
			ReadPermission:    HackMDClient.StringToNotePermissionRole(readPermission),
			WritePermission:   HackMDClient.StringToNotePermissionRole(writePermission),
			CommentPermission: HackMDClient.StringToCommentPermissionType(commentPermission),
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(note.ID, note.Title)
	},
}

var notesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a note",
	Long:  `Update a note. The note will be updated with the title provided. If no title is provided, a random title will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteId, _ := cmd.Flags().GetString("noteId")
		content, _ := cmd.Flags().GetString("content")

		if noteId == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		// TODO: add permission fields and validation

		err := api.UpdateNote(noteId, &HackMDClient.UpdateNoteOptions{
			Content: content,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Note updated")
	},
}

var notesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a note",
	Long:  `Delete a note.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteId, _ := cmd.Flags().GetString("noteId")

		if noteId == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		err := api.DeleteNote(noteId)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Note deleted")
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)

	flags.AddCommandFlags(notesCmd, []flags.FlagData{flags.OutputFlag})

	notesCmd.AddCommand(notesCreateCmd)
	flags.AddCommandFlags(notesCreateCmd, []flags.FlagData{flags.TitleFlag, flags.ContentFlag, flags.ReadPermissionFlag, flags.WritePermissionFlag, flags.CommentPermissionFlag})

	notesCmd.AddCommand(notesUpdateCmd)
	flags.AddCommandFlags(notesUpdateCmd, []flags.FlagData{flags.NoteIDFlag, flags.ContentFlag, flags.ReadPermissionFlag, flags.WritePermissionFlag, flags.CommentPermissionFlag})

	notesCmd.AddCommand(notesDeleteCmd)
	flags.AddCommandFlags(notesDeleteCmd, []flags.FlagData{flags.NoteIDFlag})
}
