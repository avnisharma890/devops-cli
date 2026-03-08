package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecIntoService(service string, shell string) error {
	fmt.Printf("Opening interactive shell in service: %s\n", service)

	cmd := exec.Command("docker", "compose", "exec", "-i", "-t", service, shell)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}