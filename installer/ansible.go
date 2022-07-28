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
var sudo = "sudo"
var apt = "apt"

func StartAnsible() bool {
	hasAnsible := checkIfExists(ansible)

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
			// err := exec.Command(python, "-m", pip, "install", "upgrade", "--user", ansible).Run()
			if errAn != nil {
				log.Fatalf("Failed to update ansible :%v", errAn)
			}
		} else {
			// Skip check if no update is required
			return true
		}
	}
	return checkIfExists(ansible)
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
	version, errVer := exec.Command(ansible, "--version").Output()
	if errVer != nil {
		log.Fatalf("Could not find ansible: %v", errVer)
	} else {
		fmt.Printf("Installed ansible successfully!\n%v\n", string(version))
	}
}

// Install ansible using python3
// 	fmt.Println("\nInstalling basic requirements to run ansible...")
// 	// TODO: add shell completion? https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html - Installing argcomplete
//
// 	// Install python
// 	hasPython := checkIfExists(python)
// 	if !hasPython {
// 		fmt.Println("\nInstalling python...")
// 		err := exec.Command(sudo, apt, "install", python).Run()
// 		if err != nil {
// 			log.Fatalf("Could not install %v: %v", python, err)
// 			return
// 		} else {
// 			fmt.Println("\nSuccessfully installed pip")
// 		}
//
// 	}
//
// 	// Install pip
// 	hasPip := checkIfExists(pip)
// 	if !hasPip {
// 		fmt.Println("\nInstalling pip...")
// 		err := exec.Command(sudo, apt, "install", "python3-pip", "-y").Run()
// 		if err != nil {
// 			log.Fatalf("Could not install %v: %v", pip, err)
// 			return
// 		} else {
// 			fmt.Println("\nSuccessfully installed pip")
// 		}
// 	}
// 	// Install ansible
// 	fmt.Println("\nInstalling ansible...")
// 	err := exec.Command(python, "-m", pip, "install", ansible, "--user").Run()
// 	if err != nil {
// 		log.Fatalf("Could not install %v: %v", ansible, err)
// 		return
// 	}
//
// 	// pip install --user ansible
//
// 	// Add .local/bin to PATH
// 	fmt.Println("\nAdding '.local/bin' to your $PATH...")
// 	path := os.Getenv("PATH")
// 	os.Setenv("PATH", "$HOME/.local/bin/"+path)
// 	fmt.Println("\nSuccessfully added '.local/bin' to your path")
//
// 	// Check ansible version
// 	version, err := exec.Command(ansible, "--version").Output()
// 	if err != nil {
// 		log.Fatalf("Could not find ansible: %v", err)
// 		fmt.Println("")
// 		return
// 	} else {
// 		fmt.Printf("Installed ansible successfully!\n%v\n", string(version))
// 	}
//
// }
