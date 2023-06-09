/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/hackmdio/hackmd-go/hackmd-cli/internal/flags"
	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/spf13/cobra"
)

// teamNotesCmd represents the teamNotes command
var teamNotesCmd = &cobra.Command{
	Use:   "teamNotes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	Short: "Create a new note",
	Long:  `Create a new note. The note will be created with the title provided. If no title is provided, a random title will be generated.`,
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
	Short: "Update a note",
	Long:  `Update a note. The note will be updated with the title provided. If no title is provided, a random title will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		content, _ := cmd.Flags().GetString("content")
		noteID, _ := cmd.Flags().GetString("noteID")
		teamPath, _ := cmd.Flags().GetString("teamPath")

		if noteID == "" {
			fmt.Println("Please provide a note ID")
			return
		}

		if teamPath == "" {
			fmt.Println("Please provide a team path")
			return
		}

		// TODO: add permission fields and validation
		err := api.UpdateTeamNote(teamPath, noteID, &HackMDClient.UpdateNoteOptions{
			Content: content,
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
	Short: "Delete a note",
	Long:  `Delete a note.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		noteID, _ := cmd.Flags().GetString("noteID")
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
