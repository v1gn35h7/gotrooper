package shell

import (
	"fmt"
	"os/exec"
)

func ExecuteScript(script string) {
	cmd := exec.Command("powershell.exe", script)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

}
