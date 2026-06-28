package cmd

import (
	"fmt"

	"github.com/sam-laister/aoc2025/services"
	"github.com/sam-laister/aoc2025/solutions"
	"github.com/spf13/cobra"
)

var day1 = &cobra.Command{
	Use:   "1",
	Short: "Day One",
	RunE: func(cmd *cobra.Command, args []string) error {
		txt, err := services.ReadDay("1")
		if err != nil {
			return err
		}

		partA, err := solutions.Day01PartA(txt)
		if err != nil {
			return err
		}

		partB, err := solutions.Day01PartB(txt)
		if err != nil {
			return err
		}

		fmt.Println("Part A:", partA)
		fmt.Println("Part B:", partB)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day1)
}
