package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for i := 1; i < 26; i++ {
		// Create modules for each day
		oldDir := "./template"
		newDir := fmt.Sprintf("./day%d", i)

		cmd := exec.Command("cp", "--recursive", oldDir, newDir)
		cmd.Run()
	}
}
