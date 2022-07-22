package installer

import (
	"automate-setup/helpers"
	"fmt"
	"log"
	"os/exec"
)

var ansible = "ansible"
var python = "python3"
var pip = "pip"

func StartAnsible() bool {
	hasAnsible := checkIfExists(ansible)

	if !hasAnsible {
		install := helpers.AskForConfirmation("Shall I install it for you dear?\nRequirements: Python3 & pip")
		if install {
			installAnsible()
		} else {
			return false
		}
	} else {
		update := helpers.AskForConfirmation("Shall I update it for you?")
		if update {
			fmt.Println("Updating ansible...")
			exec.Command(python, "-m", pip, "install", "upgrade", "--user", ansible).Run()
		} else {
			// Skip check if no update is required
			return true
		}
	}
	return checkIfExists(ansible)
}

func installAnsible() {
	fmt.Println("Installing ansible...")
	// TODO: add shell completion? https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html - Installing argcomplete
	hasPython := checkIfExists(python)
	if !hasPython {
		fmt.Println("Installing python")
		err := exec.Command(python, "--version").Run()
		if err != nil {
			log.Fatalf("Could not install python")
			return
		}
	}

	hasPip := checkIfExists(pip)
	if !hasPip {
		fmt.Println("Installing pip")
		err := exec.Command(python, "-m", pip, "-V").Run()
		if err != nil {
			log.Fatalf("Could not install pip")
			return
		}
	}

	err := exec.Command(python, "-m", pip, "install", "--user", ansible).Run()
	if err != nil {
		log.Fatalf("Could not install ansible")
		return
	}

	version, err := exec.Command(ansible, "--version").Output()
	if err != nil {
		log.Fatalf("Could not find ansible")
		return
	}

	fmt.Printf("Installed ansible successfully!\n%v\n", string(version))
}
