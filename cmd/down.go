package cmd

import (
	"fmt"
	"stackctl/internal/docker"

	"github.com/spf13/cobra"
)

var clean bool

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop the development stack",
	
	Run: func(cmd *cobra.Command, args []string) {
		err := docker.ComposeDown(clean)

		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	downCmd.Flags().BoolVar(&clean, "clean", false, "Remove orphan containers")

	rootCmd.AddCommand(downCmd)
}
