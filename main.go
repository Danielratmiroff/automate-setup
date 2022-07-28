package main

import (
	playbooks "automate-setup/execute"
	"automate-setup/helpers"
	"automate-setup/installer"
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	helpers.HandleError(err, "No user found")
	fmt.Printf("Hello %s!\n",
		user.Username)

	if user.Username != "daniel" {
		fmt.Println("\nUsername needs to be 'daniel' in order for the installation to work\nPlease update it and try again.")
		return
	}

	startAnsible := installer.StartAnsible()

	if startAnsible {
		playbooks.Start()
	}
}
