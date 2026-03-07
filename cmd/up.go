package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the development stack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting development stack...")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}