package uptime

import (
	"os/exec"
)

func Uptime() string {

	command := "uptime"
	output, error := exec.Command(command).Output()
	if error != nil {
		return ""
	} else {
		return string(output)
	}
}
