package shell

import (
	"os/exec"
	"fmt"
)

func RunCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
