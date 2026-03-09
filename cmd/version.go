package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show stackctl version",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("stackctl %s\n", Version)

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}