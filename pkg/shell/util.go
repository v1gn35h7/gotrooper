package shell

import (
	"context"
	"os/exec"
	"time"
)

func ExecuteScript(script string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if stdout, err := exec.CommandContext(ctx, "powershell.exe", script).Output(); err != nil {
		return "", err
	} else {
		return string(stdout), nil
	}
}
