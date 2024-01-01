package utils

import (
	"strings"
)

const OS_SYSTEM_INFO_CMD = "SYSTEMINFO /FO LIST"

func GetOs(platform string) string {
	var osName string
	if platform == "windows" {
		out := ExecuteCmd(OS_SYSTEM_INFO_CMD)
		list := strings.Split(out, "\n")
		if len(list) > 1 {
			os := strings.Split(list[2], ":")
			osName = strings.TrimSpace(os[1])
		}
	}

	return osName
}
