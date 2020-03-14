package services

import (
	"fmt"
	"os/exec"
)

func RunCommands(destination string, commands []string) error {
	for _, command := range commands {
		fmt.Println(fmt.Sprintf("Running: %s", command))
		commandParsed := fmt.Sprintf("cd %s && %s", destination, command)
		cmd := exec.Command("bash", "-c", commandParsed)
		//stdout, err := cmd.Output()
		_, err := cmd.Output()
		if err != nil {
			return err
		}

		//fmt.Println(stdout)
	}

	return nil
}
