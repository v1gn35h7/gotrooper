package utils

import (
	"os/exec"
)

func ExecuteCmd(cmd string) string {
	cmdCursor := exec.Command("powershell.exe", cmd)

	stdout, err := cmdCursor.Output()

	if err != nil {
		return ""
	}

	return string(stdout)
}
