package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc2025",
	Short: "Runs each aoc day",
	Long:  "Runs each oac day",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(dayOne)
	rootCmd.AddCommand(dayTwo)
}
