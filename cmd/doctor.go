package cmd

import (
	"fmt"
	"stackctl/internal/doctor"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running doctor checks...")
		doctor.CheckGo()
		doctor.CheckDocker()
		doctor.CheckDaemon()
		doctor.CheckEnvFile()
		doctor.CheckDockerCompose()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
