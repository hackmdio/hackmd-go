// build with GPT-4
// https://chat.openai.com/share/efb0f512-b21a-43f4-90be-906372bef924
package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v2"
)

func printTable(data *[]interface{}, attributes []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// set the table header
	header := make(table.Row, len(attributes))
	for i, attr := range attributes {
		header[i] = attr
	}
	t.AppendHeader(header)

	// iterate through each data item
	for _, item := range *data {
		itemValue := reflect.ValueOf(item)

		// prepare a row for this item
		row := make(table.Row, len(attributes))

		// for each attribute, find the corresponding value and append to the row
		for i, attr := range attributes {
			attrValue := itemValue.FieldByName(attr).Interface()
			row[i] = attrValue
		}

		// append the row to the table
		t.AppendRow(row)
	}

	t.Render()
}

func printNotesTable(notes *[]HackMDClient.Note) {
	// Convert slice of Note to slice of interface{}
	data := make([]interface{}, len(*notes))
	for i, v := range *notes {
		data[i] = v
	}

	// Specify the attributes you want to print
	attributes := []string{"ID", "Title"}

	printTable(&data, attributes)
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

func printCSV(data *[]interface{}, attributes []string) {
	// create a new CSV writer
	writer := csv.NewWriter(os.Stdout)

	// write the header
	err := writer.Write(attributes)
	if err != nil {
		fmt.Println("Failed to write to CSV:", err)
		return
	}

	// iterate through each item and write the attribute values
	for _, item := range *data {
		itemValue := reflect.ValueOf(item)

		// prepare a row for this item
		row := make([]string, len(attributes))

		// for each attribute, find the corresponding value and add to the row
		for i, attr := range attributes {
			attrValue := itemValue.FieldByName(attr).Interface()
			row[i] = fmt.Sprintf("%v", attrValue)
		}

		// write the row to the CSV file
		err := writer.Write(row)
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

func printNotesCSV(notes *[]HackMDClient.Note) {
	data := make([]interface{}, len(*notes))
	for i, v := range *notes {
		data[i] = v
	}
	attributes := []string{"ID", "Title"}
	printCSV(&data, attributes)
}

func PrintNotes(output string, notes *[]HackMDClient.Note) {
	switch output {
	case "table":
		printNotesTable(notes)
	case "json":
		printJSON(notes)
	case "yaml":
		printYAML(notes)
	case "csv":
		printNotesCSV(notes)
	default:
		fmt.Println("Invalid output format. Please choose from table, json, yaml, csv.")
	}
}
