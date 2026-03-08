package docker

import (
	"fmt"
	"os/exec"
)

func RestartService(service string) error {

	fmt.Printf("Restarting services : %s\n", service)

	cmd := exec.Command("docker", "compose", "restart", service)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to restart services: %s", string(output))
	}

	fmt.Printf("Services %s restarted successfully\n", service)
	return nil
}