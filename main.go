package main

import (
	"automate-setup/executor"
	"fmt"
	"io"

	// "log"

	"os/exec"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!",
		user.Username)

	// start := hasAnsible(os.Stdin, os.Stdout)
	// if start {
	executor.Exec("pkgs-playbook.yml")
	// }
}

const PROMPT = ">>"

// TODO: divide this to another file

func checkIfExists(program string) bool {
	fmt.Printf("Checking for %v\n", program)
	path, err := exec.LookPath(program)
	if err != nil {
		fmt.Printf("Didn't find '%v' executable\n", program)
		return false
	} else {
		fmt.Printf("'%v' executable found in '%s'\n", program, path)
		return true
	}
}

func hasAnsible(in io.Reader, out io.Writer) bool {
	fmt.Println("Installing ansible...")
	// scanner := bufio.NewScanner(in)

	// for {
	//fmt.Print(PROMPT)
	//scanned := scanner.Scan()
	//if !scanned {
	//  return
	//}
	// TODO: io.WriteString("") do you want to install blablblaal...
	// TODO: Update system (create file that runs apt-get update)
	// TODO: add shell completion? https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html - Installing argcomplete
	// input := scanner.Text()

	// io.WriteString(out, input)
	python := "python3"
	pip := "pip"
	ansible := "ansible"

	hasPython := checkIfExists(python)
	hasPip := checkIfExists(pip)
	hasAnsible := checkIfExists(ansible)

	if !hasPython {
		err := exec.Command(python, "--version").Run()
		if err != nil {
			fmt.Printf("Could not install %v: %v,", python, err)
			return false
		}
	}

	if !hasPip {
		err := exec.Command(python, "-m", pip, "-V").Run()
		if err != nil {
			fmt.Printf("Could not install %v: %v,", pip, err)
			return false
		}
	}

	if !hasAnsible {
		err := exec.Command(python, "-m", pip, "install", "--user", ansible).Run()
		if err != nil {
			fmt.Printf("Could not install %v: %v,", pip, err)
			return false
		}
	} else {
		fmt.Println("Updating ansible then...")
		exec.Command(python, "-m", pip, "install", "upgrade", "--user", ansible).Run()
	}

	version, err := exec.Command(ansible, "--version").Output()
	if err != nil {
		fmt.Printf("Ansible not found: %v,", err)
		return false
	} else {
		fmt.Println(string(version))
		return true
	}

}
