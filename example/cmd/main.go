package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	s := []string{"cmd.exe", "/C", "start", `ex.exe`}

	cmd := exec.Command(s[0], s[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
	os.Exit(0)
}
