package docker

import (
	"os"
	"os/exec"
)

func StreamLogs(service string) error {
	cmd := exec.Command("docker", "compose", "logs", "-f", service)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}