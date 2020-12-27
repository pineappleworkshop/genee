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
		stdout, err := cmd.CombinedOutput()
		// _, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(stdout))
			return err
		}

		// fmt.Println(string(stdout))
	}

	return nil
}
