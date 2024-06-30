package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0]) //nolint:gosec
	command.Args = cmd
	command.Stdin, command.Stdout, command.Stderr = os.Stdin, os.Stdout, os.Stderr

	command.Env = os.Environ()
	var exitErr *exec.ExitError

	for k, v := range env {
		os.Unsetenv(k)
		command.Env = append(command.Env, fmt.Sprint(k+"="+v.Value))
	}

	if err := command.Run(); err != nil {
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
	}

	return 0
}
