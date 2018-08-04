package packages

import (
	"os/exec"
	"os"
	)

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
