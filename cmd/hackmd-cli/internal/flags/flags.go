package flags

import (
	"github.com/spf13/cobra"
)

// FlagData defines the structure to store flag details.
type FlagData struct {
	Name         string
	DefaultValue string
	Description  string
}

// AddCommandFlags adds flags to the provided command.
func AddCommandFlags(cmd *cobra.Command, flagsData []FlagData) {
	for _, flagData := range flagsData {
		cmd.Flags().String(flagData.Name, flagData.DefaultValue, flagData.Description)
	}
}

func AddCommandPersistentFlags(cmd *cobra.Command, flagsData []FlagData) {
	for _, flagData := range flagsData {
		cmd.PersistentFlags().String(flagData.Name, flagData.DefaultValue, flagData.Description)
	}
}

// OutputFlag is for choosing output format.
var OutputFlag = FlagData{
	Name:         "output",
	DefaultValue: "table",
	Description:  "The output format to use. Valid options are table, json, yaml, csv",
}

// TeamPathFlag is for specifying team path.
var TeamPathFlag = FlagData{
	Name:         "teamPath",
	DefaultValue: "",
	Description:  "Team path",
}

// TitleFlag is for note title.
var TitleFlag = FlagData{
	Name:         "title",
	DefaultValue: "",
	Description:  "Title of the note",
}

// ContentFlag is for note content.
var ContentFlag = FlagData{
	Name:         "content",
	DefaultValue: "",
	Description:  "Content of the note",
}

// NoteIDFlag is for note id.
var NoteIDFlag = FlagData{
	Name:         "noteId",
	DefaultValue: "",
	Description:  "ID of the note",
}

var ReadPermissionFlag = FlagData{
	Name:         "readPermission",
	DefaultValue: "",
	Description:  "set note permission: owner, signed_in, guest",
}

var WritePermissionFlag = FlagData{
  Name:         "writePermission",
  DefaultValue: "",
  Description:  "set note permission: owner, signed_in, guest",
}

var CommentPermissionFlag = FlagData{
	Name:         "commentPermission",
	DefaultValue: "",
	Description:  "set comment permission: disabled, forbidden, owners, signed_in_users, everyone",
}
