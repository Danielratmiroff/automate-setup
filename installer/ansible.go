package installer

import (
	"automate-setup/helpers"
	"fmt"
	"log"
	"os/exec"
)

var sudo = "sudo"
var apt = "apt"

func StartAnsible() bool {
	hasAnsible := checkIfExists("ansible")

	if !hasAnsible {
		install := helpers.AskForConfirmation("Shall I install it for you dear?")

		if install {
			installAnsible()
		} else {
			return false
		}

	} else {
		update := helpers.AskForConfirmation("Shall I update it for you?")

		if update {
			fmt.Println("Updating ansible...")

			errAn := exec.Command(sudo, apt, "update", "ansible").Run()
			helpers.HandleError(errAn, "Failed to update ansible")

		} else {
			// Skip check if no update is required
			return true
		}
	}
	return checkIfExists("ansible")
}

func installAnsible() {
	fmt.Println("\nUpdating the system...")
	err := exec.Command(sudo, apt, "update").Run()
	helpers.HandleError(err, "Failed to update the system")

	fmt.Println("\nInstalling required packages...")
	errPkg := exec.Command(sudo, apt, "install", "software-properties-common").Run()
	helpers.HandleError(errPkg, "Failed to install required packages")

	fmt.Println("\nInstalling ansible's PPA repository...")
	errRp := exec.Command(sudo, "add-apt-repository", "--yes", "--update", "ppa:ansible/ansible").Run()
	helpers.HandleError(errRp, "Failed to add required repository")

	fmt.Println("\nInstalling ansible...")
	errAn := exec.Command(sudo, apt, "install", "ansible", "-y").Run()
	helpers.HandleError(errAn, "Failed to install ansible")

	// Check ansible version
	version, errVer := exec.Command("ansible", "--version").Output()
	if errVer != nil {
		log.Fatalf("Could not find ansible: %v", errVer)
	} else {
		fmt.Printf("Installed ansible successfully!\n%v\n", string(version))
	}
}
