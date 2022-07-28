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
	sudoPass := helpers.AskForInput("\nPlease enter your sudo/admin password\n")
	if sudoPass == "" {
		fmt.Println("Sorry, we need those juicy admin rights")
		return
	}

	setup := &Playbook{
		name:  "setup-playbook",
		print: true,
	}
	RunPlaybook(setup, sudoPass, false)

	pkgs := &Playbook{
		name:  "play-pkgs",
		print: true,
	}
	RunPlaybook(pkgs, sudoPass, false)

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
	// zshrc := &Playbook{
	// 	name:  "play-zshrc",
	// 	print: false,
	// }
	fish := &Playbook{
		name:  "play-fish",
		print: false,
	}
	secondaryPkgs := &Playbook{
		name:  "play-secondary-pkgs",
		print: false,
	}

	// continue here :: need to run once more docker-run to test
	// TODO: refactor playbook struct, waitgroup bool and printing all playbooks
	wg.Add(4)
	go RunPlaybook(secondaryPkgs, sudoPass, true)
	go RunPlaybook(docker, sudoPass, true)
	go RunPlaybook(lazygit, sudoPass, true)
	go RunPlaybook(neovim, sudoPass, true)
	wg.Wait()

	// RunPlaybook(zshrc, sudoPass, true)
	RunPlaybook(fish, sudoPass, true)

	dotfiles := &Playbook{
		name:  "play-dotfiles",
		print: true,
	}
	RunPlaybook(dotfiles, sudoPass, false)

	// mongo := &Playbook{
	//  name:  "play-mongo",
	//  print: false,
	//}
	// RunPlaybook(mongodb, sudoPass) -- disabled bc playbook is unfinished (maybe unnecessary?)
}

func RunPlaybook(playbook *Playbook, sudoPass string, wait_group bool) {
	ansibleFolder := "ansible/"
	inventoryPath := ansibleFolder + "inventory.yml"
	formattedFile := ansibleFolder + playbook.name + ".yml"
	becomeSudo := "-e ansible_become_pass=" + sudoPass

	// cmd := exec.Command("ansible-playbook", "--diff", "-i", inventoryPath, formattedFile, becomeSudo, "--check") // check for testing
	cmd := exec.Command("ansible-playbook", "-i", inventoryPath, formattedFile, becomeSudo)

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
	if wait_group {
		defer wg.Done()
	}
}
