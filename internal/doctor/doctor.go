package doctor

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckDocker() {
	// check if docker is installed
	cmd := exec.Command("docker", "--version")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Docker not installed!")
		return
	}

	fmt.Println("Docker detected: ", string(output))
}

func CheckGo() {
	// check if go is installed
	cmd := exec.Command("go", "version")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Go not installed!")
		return
	}

	fmt.Println("Go detected: ", string(output))
}

func CheckDaemon() {
	// check if docker daemon is running
	cmd := exec.Command("docker", "info")

	output, error := cmd.Output()

	if error != nil {
		fmt.Println("Docker daemon not running")
		return
	}

	fmt.Println("Docker daemon running", string(output))
}

func CheckEnvFile() {
	// check if .env file exists
	_, error := os.Stat(".env")

	if error != nil {
		fmt.Println(".env file not found")
		return
	}

	fmt.Println(".env file detected")
}

func CheckDockerCompose() {
	cmd := exec.Command("docker", "compose", "version")

	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Docker compose not installed!")
		return
	}

	fmt.Println("Docker compose detected")
}