/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/hackmdio/hackmd-go/hackmd-cli/internal"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	historyCmd.Flags().String("output", "table", "The output format to use. Valid options are table, json, yaml, csv")
}
