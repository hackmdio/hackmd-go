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
	Short: "List all notes of the user",
	Long:  `The 'notes' command retrieves and lists all the notes associated with the user's HackMD account. It fetches the note list and provides an organized output which can be optionally formatted based on user's preference.`,
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
	Long:  `The 'create' command initiates the creation of a new note in the user's HackMD account. The user can specify title and content as well as the permissions for reading, writing, and commenting. If no content is provided, it takes standard input. The note creation might fail in case of errors, and the system handles them gracefully.`,
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
	Short: "Update an existing note",
	Long:  `The 'update' command allows users to modify an existing note in their HackMD account. The user can change the content of the note and adjust the permissions for reading, writing, and commenting. The command requires a note ID, and returns an error if no ID is provided or if any issues occur during the update process.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteId, _ := cmd.Flags().GetString("noteId")
		content, _ := cmd.Flags().GetString("content")
		readPermission, _ := cmd.Flags().GetString("readPermission")
		writePermission, _ := cmd.Flags().GetString("writePermission")
		commentPermission, _ := cmd.Flags().GetString("commentPermission")

		if noteId == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		err := api.UpdateNote(noteId, &HackMDClient.UpdateNoteOptions{
			Content:           content,
			ReadPermission:    HackMDClient.StringToNotePermissionRole(readPermission),
			WritePermission:   HackMDClient.StringToNotePermissionRole(writePermission),
			CommentPermission: HackMDClient.StringToCommentPermissionType(commentPermission),
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
	Short: "Delete a specific note",
	Long:  `The 'delete' command allows users to remove a specific note from their HackMD account. The note is identified by its ID. An error message is displayed if a note ID is not provided or if any problems occur during the deletion process.`,
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
