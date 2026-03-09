package cmd

import (
	"fmt"
	"stackctl/internal/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show loaded stack configuration",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Loading stack configuration...")

		cfg, err := config.LoadConfig("stackctl.yaml")

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Configured services:")

		for name, service := range cfg.Services {
			fmt.Printf("- %s running on port %d\n", name, service.Port)
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}