package execute

import (
	"automate-setup/helpers"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Start(username string) {

	installAll := helpers.AskForConfirmation("\nDo you want to run all ansible playbooks?")
	if !installAll {
		fmt.Println("All righty!, bye now.")
		return
	}

	sudoPass := helpers.AskForInput("\nPlease enter your sudo/admin password\n")
	if sudoPass == "" {
		fmt.Println("Sorry, we need those juicy admin rights")
		return
	}

	files, err := ioutil.ReadDir("ansible/")
	helpers.HandleError(err, "Couldn't find 'ansible/' folder")

	for _, file := range files {
		fileName := file.Name()
		if strings.Contains(fileName, "playbook") && !file.IsDir() {
			RunPlaybook(fileName, sudoPass)
		}
	}
}

func RunPlaybook(playbook string, sudoPass string) {
	ansibleFolder := "ansible/"
	inventoryPath := ansibleFolder + "inventory.yml"
	formattedFile := ansibleFolder + playbook + ".yml"
	becomeSudo := "-e ansible_become_pass=" + sudoPass

	cmd := exec.Command("ansible-playbook", "-i", inventoryPath, formattedFile, becomeSudo)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	helpers.HandleError(err, "cmd.Run() failed with %s\n")

	// Hidden outStr
	_, errStr := stdoutBuf.String(), stderrBuf.String()
	if errStr != "" {
		fmt.Printf("err:\n%s\n", errStr)
	}
}
