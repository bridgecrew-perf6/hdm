package exec

import (
	"fmt"
	"log"
	"os"
	osexec "os/exec"
	"strings"
)

func ExecuteCommand(execCmd string) {
	fmt.Println(execCmd)
	splitCmd := strings.Split(execCmd, " ")
	cmdName := splitCmd[0]
	cmdArgs := splitCmd[1:]

	command := osexec.Command(cmdName, cmdArgs...)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
