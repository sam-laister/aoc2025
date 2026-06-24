package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02PartA(input string, verbose bool) (int, error) {
	pairs := strings.Split(strings.TrimSpace(input), ",")

	var invalidIds []int

	for _, i := range pairs {
		pair := strings.Split(i, "-")

		lowerBound, _ := strconv.Atoi(pair[0])
		upperBound, _ := strconv.Atoi(pair[1])

	main:
		for diff := range upperBound - lowerBound + 1 {
			curr := diff + lowerBound
			strValue := strconv.Itoa(curr)
			length := len(strValue)

			if length%2 != 0 {
				continue
			}

			midpoint := (length / 2)

			if verbose {
				fmt.Printf("LB: %d, UB: %d, M: %d\n", lowerBound, upperBound, midpoint)
			}

			for index := range midpoint {
				if strValue[index] != strValue[midpoint+index] {
					continue main
				}
			}

			if verbose {
				fmt.Printf("Found match: %d\n", curr)
			}

			invalidIds = append(invalidIds, curr)
		}
	}

	total := 0
	for _, num := range invalidIds {
		total += num
	}

	return total, nil
}
