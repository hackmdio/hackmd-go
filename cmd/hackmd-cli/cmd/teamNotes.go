/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/cmd/hackmd-cli/internal"
	"github.com/hackmdio/hackmd-go/cmd/hackmd-cli/internal/flags"
	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/spf13/cobra"
)

// teamNotesCmd represents the teamNotes command
var teamNotesCmd = &cobra.Command{
	Use:   "teamNotes",
	Short: "List all notes of a specific team",
	Long:  `The 'teamNotes' command retrieves and lists all the notes associated with a specified team in the user's HackMD account. The user needs to provide the team path for the team whose notes need to be fetched. If a team path isn't provided, the command will return an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		teamPath, _ := cmd.Flags().GetString("teamPath")
		api := internal.GetHackMDClient()

		if teamPath == "" {
			fmt.Println("Please provide a team path")
			return
		}

		notes, err := api.GetTeamNotes(teamPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		output := cmd.Flag("output").Value.String()
		internal.PrintNotes(output, notes)
	},
}

var teamNotesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note for a team",
	Long:  `The 'create' command initiates the creation of a new note for a specific team in the user's HackMD account. The user can specify title, content and permissions, as well as the team path. If a team path or content isn't provided, the command will return an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		title, _ := cmd.Flags().GetString("title")
		content, _ := cmd.Flags().GetString("content")
		teamPath, _ := cmd.Flags().GetString("teamPath")
		readPermission, _ := cmd.Flags().GetString("readPermission")
		writePermission, _ := cmd.Flags().GetString("writePermission")
		commentPermission, _ := cmd.Flags().GetString("commentPermission")

		if teamPath == "" {
			fmt.Println("Please provide a team path")
			return
		}

		note, err := api.CreateTeamNote(teamPath, &HackMDClient.CreateNoteOptions{
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

var teamNotesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a note for a team",
	Long:  `The 'update' command allows users to modify an existing note in a specific team in their HackMD account. The user can adjust the content and permissions for the note, and needs to provide the note ID and team path. If these aren't provided, the command will return an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		content, _ := cmd.Flags().GetString("content")
		noteID, _ := cmd.Flags().GetString("noteId")
		teamPath, _ := cmd.Flags().GetString("teamPath")
		readPermission, _ := cmd.Flags().GetString("readPermission")
		writePermission, _ := cmd.Flags().GetString("writePermission")
		commentPermission, _ := cmd.Flags().GetString("commentPermission")

		if noteID == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		if teamPath == "" {
			fmt.Println("Please provide a team path")
			return
		}

		err := api.UpdateTeamNote(teamPath, noteID, &HackMDClient.UpdateNoteOptions{
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

var teamNotesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a note from a team",
	Long:  `The 'delete' command allows users to remove a specific note from a specified team in their HackMD account. The note and team are identified by their respective IDs. An error message is displayed if the note ID or team path is not provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteID, _ := cmd.Flags().GetString("noteId")
		teamPath, _ := cmd.Flags().GetString("teamPath")

		if noteID == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		if teamPath == "" {
			fmt.Println("Please provide a team path")
			return
		}

		err := api.DeleteTeamNote(teamPath, noteID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Note deleted")
	},
}

func init() {
	rootCmd.AddCommand(teamNotesCmd)
	flags.AddCommandFlags(teamNotesCmd, []flags.FlagData{flags.OutputFlag})
	flags.AddCommandPersistentFlags(teamNotesCmd, []flags.FlagData{flags.TeamPathFlag})

	teamNotesCmd.AddCommand(teamNotesCreateCmd)
	flags.AddCommandFlags(teamNotesCreateCmd, []flags.FlagData{flags.TitleFlag, flags.ContentFlag, flags.ReadPermissionFlag, flags.WritePermissionFlag, flags.CommentPermissionFlag})

	teamNotesCmd.AddCommand(teamNotesUpdateCmd)
	flags.AddCommandFlags(teamNotesUpdateCmd, []flags.FlagData{flags.NoteIDFlag, flags.ContentFlag, flags.ReadPermissionFlag, flags.WritePermissionFlag, flags.CommentPermissionFlag})

	teamNotesCmd.AddCommand(teamNotesDeleteCmd)
	flags.AddCommandFlags(teamNotesDeleteCmd, []flags.FlagData{flags.NoteIDFlag})
}
