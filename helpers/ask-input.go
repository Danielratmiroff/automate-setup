package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func AskForInput(s string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(s)

		response, err := reader.ReadString('\n')
		HandleError(err, "Error retrieving input")

		return response
	}
}
