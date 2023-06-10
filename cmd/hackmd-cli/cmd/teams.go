package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/hackmdio/hackmd-go/hackmd-cli/internal/flags"
	"github.com/spf13/cobra"
)

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "List all teams",
	Long:  `The 'teams' command retrieves and lists all the teams associated with the user's HackMD account. The command requires an active user session. If not logged in, it will return an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()
		output := cmd.Flag("output").Value.String()

		teams, err := api.GetTeams()
		if err != nil {
			fmt.Println(err)
			return
		}

		internal.PrintTeams(output, teams)
	},
}

func init() {
	rootCmd.AddCommand(teamsCmd)
	flags.AddCommandFlags(teamsCmd, []flags.FlagData{flags.OutputFlag})
}
