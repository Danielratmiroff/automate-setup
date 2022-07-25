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
var apt = "apt-get"

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
			err := exec.Command(python, "-m", pip, "install", "upgrade", "--user", ansible).Run()
			if err != nil {
				log.Fatalf("Failed to update ansible :%v", err)
			}
		} else {
			// Skip check if no update is required
			return true
		}
	}
	return checkIfExists(ansible)
}

func installAnsible() {
	fmt.Println("\nInstalling ansible...")
	// TODO: add shell completion? https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html - Installing argcomplete
	hasPython := checkIfExists(python)
	if !hasPython {
		fmt.Println("\nInstalling python...")
		err := exec.Command(apt, "install", python).Run()
		if err != nil {
			log.Fatalf("Could not install python: %v", err)
			return
		}
	}

	hasPip := checkIfExists(pip)
	if !hasPip {
		fmt.Println("\nInstalling pip...")
		err := exec.Command(apt, "install", "python3-pip", "-y").Run()
		if err != nil {
			log.Fatalf("Could not install pip: %v", err)
			return
		}
	}

	fmt.Println("\nInstalling ansible...")
	err := exec.Command(python, "-m", pip, "install", ansible).Run()
	if err != nil {
		log.Fatalf("Could not install ansible: %v", err)
		return
	}

	version, err := exec.Command(ansible, "--version").Output()
	if err != nil {
		log.Fatalf("Could not find ansible: %v", err)
		return
	}

	fmt.Printf("Installed ansible successfully!\n%v\n", string(version))
}
