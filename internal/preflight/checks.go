package preflight

import (
	"fmt"
	"os"
	"os/exec"
)

func RunChecks() error {

	fmt.Println("Running pre-deployment checks...")

	// Check Docker CLI
	if _, err := exec.LookPath("docker"); err != nil {
		return fmt.Errorf("docker is not installed")
	}
	fmt.Println("✔ Docker CLI found")

	// Check Docker daemon
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker daemon is not running")
	}
	fmt.Println("✔ Docker daemon running")

	// Check compose file
	if _, err := os.Stat("docker-compose.yml"); err != nil {
		return fmt.Errorf("docker-compose.yml not found")
	}
	fmt.Println("✔ docker-compose.yml detected")

	// Check stackctl config
	if _, err := os.Stat("stackctl.yaml"); err != nil {
		fmt.Println("⚠ stackctl.yaml not found (optional)")
	} else {
		fmt.Println("✔ stackctl.yaml detected")
	}

	fmt.Println("Environment checks passed.\n")

	return nil
}