package internal;

import (
	"os"
  "fmt"
  "encoding/json"
	"github.com/jedib0t/go-pretty/v6/table"
  HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
)

func printTable(notes *[]HackMDClient.Note) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title"})
	for _, note := range *notes {
		t.AppendRow(table.Row{note.ID, note.Title})
	}
	t.Render()
}

func printJSON(notes *[]HackMDClient.Note) {
	jsonNotes, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		fmt.Println("Failed to convert to JSON:", err)
		return
	}
	fmt.Println(string(jsonNotes))
}


func PrintNotes(output string, notes *[]HackMDClient.Note) {
	switch output {
	case "table":
		printTable(notes)
	case "json":
		printJSON(notes)
	case "yaml":
		// TODO: Implement YAML formatting
	case "csv":
		// TODO: Implement CSV formatting
	default:
		fmt.Println("Invalid output format. Please choose from table, json, yaml, csv.")
	}
}
