package docker

import (
	"fmt"
	"os/exec"
	"stackctl/internal/preflight"
)

func Deploy() error {
	fmt.Println("Starting deployment pipeline...")

	if err := preflight.RunChecks(); err != nil {
		return err
	}

	fmt.Println("Step 1: Building Docker image...")

	buildCommand := exec.Command("docker", "compose", "build")

	output, err := buildCommand.CombinedOutput()

	if err != nil {
		return fmt.Errorf("build failed: %s", string(output))
	}

	fmt.Println("Image built successfully.")

	fmt.Println("Step 2: Restarting services...")

	upCmd := exec.Command("docker", "compose", "up", "-d")

	output, err = upCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("service restart failed: %s", output)
	}

	fmt.Println("Deployment completed successfully.")

	return nil
}