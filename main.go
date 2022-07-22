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

	hasAnsible := installer.StartAnsible()

	if hasAnsible {
		playbooks.Start()
	}
}
