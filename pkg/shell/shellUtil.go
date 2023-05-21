package shell

import (
	"os/exec"
)

func ExecuteScript(script string) (string, error) {
	cmd := exec.Command("powershell.exe", script)

	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
