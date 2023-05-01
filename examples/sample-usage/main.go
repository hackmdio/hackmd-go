package main

import (
	"os"
	"fmt"
	HackMDClient "github.com/hackmdio/go-api-client/pkg/api"
)

func main() {
	// get access token and endpoint from environment variable
	var accessToken = os.Getenv("HMD_ACCESS_TOKEN")
	var apiEndpoint = os.Getenv("HMD_API_ENDPOINT")

	if accessToken == "" {
	  // panic if access token is not set
	  panic("Missing access token")
	}

	var client *HackMDClient.APIClient

	if apiEndpoint == "" {
    // use default endpoint if not set
    client = HackMDClient.NewAPIClient(accessToken)
  } else {
	  client = HackMDClient.NewAPIClient(accessToken, HackMDClient.WithAPIEndpointURL(apiEndpoint))
	}

	var user, err = client.GetMe()

	if err != nil {
		panic(err)
	}

	println("Hello, " + user.Name)

	var notes, err2 = client.GetHistory()

	if err != nil {
		panic(err2)
	}

	println(fmt.Sprintf("History length: %d", len(*notes)))

	// get one note from history
	var noteID = (*notes)[0].ID
	var note2, err3 = client.GetNote(noteID)

	if err3 != nil {
	  panic(err3)
	}

	println(fmt.Sprintf("Note 2 title: %s", note2.Title))
	println(fmt.Sprintf("Note 2 content: %s", note2.Content))

	// create a new note
	var note, err4 = client.CreateNote(&HackMDClient.CreateNoteOptions{ Title: "Hello, world!" })

  if err4 != nil {
    panic(err4)
  }

  println(fmt.Sprintf("Note ID: %s", note.ID))

  // update the note
  var err5 = client.UpdateNote(note.ID, &HackMDClient.UpdateNoteOptions{ Content: "# Hello, world! (updated)" })

  if err5 != nil {
    panic(err5)
  }

  // delete the note
  var err6 = client.DeleteNote(note.ID)

  if err6 != nil {
    panic(err6)
  }
}
