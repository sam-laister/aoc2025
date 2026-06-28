package cmd

import (
	"fmt"

	"github.com/sam-laister/aoc2025/services"
	"github.com/sam-laister/aoc2025/solutions"
	"github.com/spf13/cobra"
)

var day4 = &cobra.Command{
	Use:   "4",
	Short: "Day 4",
	RunE: func(cmd *cobra.Command, args []string) error {
		txt, err := services.ReadDay("4")
		if err != nil {
			return err
		}

		partA, err := solutions.Day04PartA(txt, false)
		if err != nil {
			return err
		}

		partB, err := solutions.Day04PartB(txt, false)
		if err != nil {
			return err
		}

		fmt.Println("Part A:", partA)
		fmt.Println("Part B:", partB)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day4)
}
