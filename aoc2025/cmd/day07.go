package cmd

import (
	"fmt"

	"github.com/sam-laister/aoc2025/services"
	"github.com/sam-laister/aoc2025/solutions"
	"github.com/spf13/cobra"
)

var day7 = &cobra.Command{
	Use:   "7",
	Short: "Day 7",
	RunE: func(cmd *cobra.Command, args []string) error {
		txt, err := services.ReadDay("7")
		if err != nil {
			return err
		}

		partA, err := solutions.Day07PartA(txt, false)
		if err != nil {
			return err
		}

		partB, err := solutions.Day07PartB(txt, false)
		if err != nil {
			return err
		}

		fmt.Println("Part A:", partA)
		fmt.Println("Part B:", partB)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day7)
}
