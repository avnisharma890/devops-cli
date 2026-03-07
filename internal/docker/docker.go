package docker

import (
	"fmt"
	"os/exec"
)

func ComposeUp() error {
	cmd := exec.Command("docker", "compose", "up", "-d")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to start services: %s", output)
	}

	fmt.Println("development stack started")
	return nil
}
