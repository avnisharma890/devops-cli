package cmd

import (
	"fmt"
	"stackctl/internal/docker"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use: "exec [service] [shell]",
	Short: "Open an interactive shell inside a container",
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		shell :=  args[1]

		fmt.Printf("Connecting to service: %s\n", service)

		err := docker.ExecIntoService(service, shell)

		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}