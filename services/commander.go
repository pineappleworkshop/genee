package services

import (
	"fmt"
	"os/exec"
)

func RunCommands(destination string, commands []string) error {
	for _, command := range commands {
		commandParsed := fmt.Sprintf("cd %s && %s", destination, command)
		err := RunCommand(commandParsed)
		if err != nil {
			return err
		}
	}

	return nil
}

func RunCommand(command string) error {
	fmt.Println(fmt.Sprintf("Running: %s", command))
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(stdout))
		return err
	}

	return nil
}
