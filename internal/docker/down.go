package docker

import (
	"fmt"
	"os/exec"
)

func ComposeDown(clean bool) error {
	args := []string{"compose", "down"}

	if clean {
		args = append(args, "--remove-orphans")
	}

	cmd := exec.Command("docker", args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to stop services: %s", output)
	}

	fmt.Println("Development stack stopped")
	return nil 
}