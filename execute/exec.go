package execute

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func Exec(command string) {

	ansible := "ansible"
	hosts := fmt.Sprintf("%v/hosts.ini", ansible)
	playbook := fmt.Sprintf("%v/%v", ansible, command)

	// TODO: this could be a string "ansible-.. -key... hosts" that gets splitted with commas
	cmd := exec.Command("ansible-playbook", "--diff", "-Ki", hosts, playbook)
	// cmd := exec.Command("ansible-playbook", "-i", ansibleHost, filepath, "--syntax-check" )

	// if runtime.GOOS == "windows" {
	// cmd = exec.Command("tasklist")
	// }

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// TODO: as of now, prints double messages (progress and end)
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

}
