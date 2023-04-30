package main

import (
  "os"
  HackMDClient "github.com/hackmdio/go-client"
)

func main() {
  // get access token from environment variable
  var client = HackMDClient.NewAPIClient(os.Getenv("HMD_ACCESS_TOKEN"))

  var user, err = client.GetMe()

  if err != nil {
    panic(err)
  }

  println("Hello " + user.Name)
}
