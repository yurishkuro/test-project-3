package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Cmd{
		Path:   "/bin/bash",
		Args:   []string{"bash", "-c", "lsof -iTCP -sTCP:LISTEN -P +c0"},
		Stdout: os.Stderr,
		Stderr: os.Stderr,
	}
	err := cmd.Start()
	println("error if any:", err)
}
