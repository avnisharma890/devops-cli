package cmd

import (
	"fmt"
	"stackctl/internal/docker"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs [service]",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]

		err := docker.StreamLogs(service)
		if err != nil {
			fmt.Println("Error streaming logs: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
