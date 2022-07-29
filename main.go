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

	username := user.Username
	fmt.Printf("Hello %s!\n", username)

	startAnsible := installer.StartAnsible()

	if startAnsible {
		playbooks.Start(username)
	}
}
