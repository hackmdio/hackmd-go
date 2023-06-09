package internal;

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func GetMeFlow(showLoginMessage bool) {
	api := GetHackMDClient()
	user, err := api.GetMe()

	if err != nil {
		fmt.Println(err)
		return
	}

	if showLoginMessage {
    fmt.Println("You are now logged in.")
  }

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("2"))

	// fmt.Printf("Logged in as: %s\nusername: %s\n", style.Render(user.Email), style.Render(user.UserPath))
	fmt.Printf("Using API Endpoint: %s\n", style.Render(api.GetHackmdAPIEndpointURL()))

	PrintTable(&[]interface{}{user}, []string{"Email", "UserPath"})
}
