package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func Day03PartA(input string, verbose bool) (int, error) {
	batteries := strings.Split(input, "\n")

	total := 0

	for _, battery := range batteries {
		if len(battery) < 2 {
			continue
		}

		highest := [2]int{0, 0}

		var batteryInt []int

		for _, c := range battery {
			bInt, err := strconv.Atoi(string(c))

			if err != nil {
				return -1, err
			}

			batteryInt = append(batteryInt, bInt)
		}

		start := -1
		for loop := range 2 {
			for index, value := range batteryInt {
				if index <= start {
					continue
				}

				if loop == 0 && index == len(batteryInt)-1 {
					continue
				}

				if value > highest[loop] {
					start = index
					highest[loop] = value
				}
			}
		}

		sum, _ := strconv.Atoi(fmt.Sprintf("%d%d", highest[0], highest[1]))

		if verbose {
			fmt.Printf("Adding %d\n", sum)
		}

		total += sum
	}

	return total, nil
}

func Day03PartB(input string, verbose bool) (int, error) {
	batteries := strings.Split(input, "\n")

	total := 0

	for _, battery := range batteries {
		if len(battery) < 2 {
			continue
		}

		highest := [12]int{}

		var batteryInt []int

		for _, c := range battery {
			bInt, err := strconv.Atoi(string(c))

			if err != nil {
				return -1, err
			}

			batteryInt = append(batteryInt, bInt)
		}

		start := -1
		for loop := range 12 {
			for index, value := range batteryInt {
				if index <= start {
					continue
				}

				if index > len(batteryInt)-12+loop {
					continue
				}

				if value > highest[loop] {
					start = index
					highest[loop] = value
				}
			}
		}

		var sb strings.Builder
		for _, v := range highest {
			sb.WriteByte(byte('0' + v))
		}
		sum, _ := strconv.Atoi(sb.String())

		if verbose {
			fmt.Printf("Adding %d\n", sum)
		}

		total += sum
	}

	return total, nil
}
