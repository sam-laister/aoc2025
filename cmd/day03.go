package cmd

import (
	"fmt"

	"github.com/sam-laister/aoc2025/services"
	"github.com/sam-laister/aoc2025/solutions"
	"github.com/spf13/cobra"
)

var day3 = &cobra.Command{
	Use:   "3",
	Short: "Day Three",
	RunE: func(cmd *cobra.Command, args []string) error {
		txt, err := services.ReadDay("3")
		if err != nil {
			return err
		}

		partA, err := solutions.Day03PartA(txt, false)
		if err != nil {
			return err
		}

		partB, err := solutions.Day03PartB(txt, false)
		if err != nil {
			return err
		}

		fmt.Println("Part A:", partA)
		fmt.Println("Part B:", partB)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day3)
}
