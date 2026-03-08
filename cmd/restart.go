package cmd

import (
	"fmt"
	"stackctl/internal/docker"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [service]",
	Short: "Restart a specific service",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		service := args[0]

		fmt.Printf("Preparing to restart service: %s\n", service)

		err := docker.RestartService(service)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Restart command completed.")
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}