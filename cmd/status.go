package cmd

import (
	"fmt"
	"stackctl/internal/docker"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of services",

	Run: func(cmd *cobra.Command, args []string) {

		err := docker.ComposeStatus()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}