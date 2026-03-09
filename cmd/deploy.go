package cmd

import (
	"fmt"
	"stackctl/internal/docker"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Run deployment pipeline",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Preparing deployment...")

		err := docker.Deploy()

		if err != nil {
			fmt.Println("Deployment failed:", err)
			return
		}

		fmt.Println("Deploy command finished.")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}