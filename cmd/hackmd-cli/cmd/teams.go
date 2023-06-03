package cmd

import (
	"fmt"
	"os"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "List all teams",
	Long: `List all teams in your HackMD account.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := internal.GetHackMDClient()

		teams, err := api.GetTeams()
		if err != nil {
			fmt.Println(err)
			return
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Path", "Name"})

		for _, team := range teams {
			t.AppendRow(table.Row{team.Path, team.Name})
		}

		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(teamsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teamsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
