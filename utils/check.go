package utils

import "os/exec"

// CheckRequiredInstalled checks if the required apps is installed first
// such as Node, NPM, PHP, etc...
// based from: https://siongui.github.io/2018/03/16/go-check-if-command-exists/
func CheckRequiredInstalled(command string, arg string) bool {
	cmd := exec.Command(command, arg)

	// returns false if the command fails
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}
