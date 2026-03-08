package docker

import (
	"fmt"
	"os/exec"
)

func ComposeStatus() error {
	cmd := exec.Command("docker", "compose", "ps")
	
	output, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("Failed to get service status")
	}

	fmt.Println(string(output))
	return nil
}
