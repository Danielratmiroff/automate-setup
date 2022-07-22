package execute

import (
	"automate-setup/helpers"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

type Playbook struct {
	name  string
	print bool
}

var wg sync.WaitGroup

func Start() {
	setup := &Playbook{
		name:  "setup-playbook",
		print: true,
	}

	sudoPass := helpers.AskForInput("Please enter your sudo/admin password")
	if sudoPass == "" {
		return
	}

	wg.Add(2)
	go RunPlaybook(setup, sudoPass)
	go RunPlaybook(setup, sudoPass)
	wg.Wait()
}

func RunPlaybook(playbook *Playbook, sudoPass string) {
	// TODO: reformat these variables -- continue here
	ansibleFolder := "ansible/"
  inventoryPath := ansibleFolder + "inventory.yaml"
	formattedFile := ansibleFolder + playbook.name  ".yml"
	adminRights := "ansible_become_pass=" + sudoPass

	cmd := exec.Command("ansible-playbook", "--diff", "-i", inventoryPath, formattedFile, "-e", adminRights, "--check")

	if playbook.print {

		var stdoutBuf, stderrBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		// Hidden outStr
		_, errStr := stdoutBuf.String(), stderrBuf.String()
		if errStr != "" {
			fmt.Printf("err:\n%s\n", errStr)
		}
	}
	defer wg.Done()
}
