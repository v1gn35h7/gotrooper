package shell

import (
	"fmt"
	"testing"
)

func TestExcuteScript(t *testing.T) {

	out, err := ExecuteScript("Get-Process")

	if err != nil {
		t.Error("Failed")
	}

	fmt.Println(out)

}
