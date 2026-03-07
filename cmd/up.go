package cmd

import (
	"fmt"
	"stackctl/internal/docker"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the development stack",
	Run: func(cmd *cobra.Command, args []string) {
		err := docker.ComposeUp()

		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}