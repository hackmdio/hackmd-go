package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v2"
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

func printYAML(notes *[]HackMDClient.Note) {
	yamlNotes, err := yaml.Marshal(notes)
	if err != nil {
		fmt.Println("Failed to convert to YAML:", err)
		return
	}
	fmt.Println(string(yamlNotes))
}

func printCSV(notes *[]HackMDClient.Note) {
	writer := csv.NewWriter(os.Stdout)

	// write header
	err := writer.Write([]string{"ID", "Title"})
	if err != nil {
		fmt.Println("Failed to write to CSV:", err)
		return
	}

	// write data
	for _, note := range *notes {
		err := writer.Write([]string{note.ID, note.Title})
		if err != nil {
			fmt.Println("Failed to write to CSV:", err)
			return
		}
	}

	writer.Flush()
	if writer.Error() != nil {
		fmt.Println("Failed to write to CSV:", writer.Error())
	}
}

func PrintNotes(output string, notes *[]HackMDClient.Note) {
	switch output {
	case "table":
		printTable(notes)
	case "json":
		printJSON(notes)
	case "yaml":
		printYAML(notes)
	case "csv":
		printCSV(notes)
	default:
		fmt.Println("Invalid output format. Please choose from table, json, yaml, csv.")
	}
}
