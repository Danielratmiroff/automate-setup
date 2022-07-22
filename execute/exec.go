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

	// TODO: ask for user name to login into

	sudoPass := helpers.AskForInput("Please enter your sudo/admin password\n")
	if sudoPass == "" {
		fmt.Println("Sorry, we need those juicy admin rights")
		return
	}

	setup := &Playbook{
		name:  "setup-playbook",
		print: true,
	}
	RunPlaybook(setup, sudoPass)

	pkgs := &Playbook{
		name:  "play-pkgs",
		print: true,
	}
	RunPlaybook(pkgs, sudoPass)

	docker := &Playbook{
		name:  "play-docker",
		print: false,
	}
	lazygit := &Playbook{
		name:  "play-lazygit",
		print: false,
	}
	neovim := &Playbook{
		name:  "play-neovim",
		print: true,
	}
	zshrc := &Playbook{
		name:  "play-zshrc",
		print: false,
	}
	wg.Add(4)
	go RunPlaybook(docker, sudoPass)
	go RunPlaybook(lazygit, sudoPass)
	go RunPlaybook(neovim, sudoPass)
	go RunPlaybook(zshrc, sudoPass)
	wg.Wait()

	dotfiles := &Playbook{
		name:  "play-dotfiles",
		print: true,
	}
	RunPlaybook(dotfiles, sudoPass)

	// mongo := &Playbook{
	//  name:  "play-mongo",
	//  print: false,
	//}
	// RunPlaybook(mongodb, sudoPass) -- disabled bc playbook is unfinished (maybe unnecessary?)
}

func RunPlaybook(playbook *Playbook, sudoPass string) {
	ansibleFolder := "ansible"
	inventoryPath := fmt.Sprintf("%v/inventory.yml", ansibleFolder)               // ansible/inventory.yml
	formattedFile := fmt.Sprintf("%v/%v%v", ansibleFolder, playbook.name, ".yml") // ansible/playbook.yml
	becomeSudo := fmt.Sprintf("-e ansible_become_pass=%v", sudoPass)              // sudo user pwd

	cmd := exec.Command("ansible-playbook", "--diff", "-i", inventoryPath, formattedFile, becomeSudo, "--check")

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
