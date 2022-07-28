package installer

import (
	"fmt"
	"os/exec"
)

func checkIfExists(program string) bool {
	fmt.Printf("Checking for %v\n", program)
	path, err := exec.LookPath(program)
	if err != nil {
		fmt.Printf("Didn't find '%v' executable\n", program)
		return false
	} else {
		fmt.Printf("Found in '%s'\n\n", path)
		return true
	}
}
