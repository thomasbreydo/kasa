package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kasa",
	Short: "A command-line utility to interact with Kasa smart devices",
	Long:  "A command-line utility to interact with Kasa smart devices.",
}

func init() {
	rootCmd.AddCommand(OnCmd)
	rootCmd.AddCommand(OffCmd)
	rootCmd.AddCommand(StatusCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
