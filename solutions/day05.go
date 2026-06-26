package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Lower int
	Upper int
}

func NewRange(input string) (Range, error) {
	values := strings.Split(input, "-")
	if len(values) != 2 {
		return Range{}, fmt.Errorf("invalid range: %q", input)
	}
	lower, err := strconv.Atoi(values[0])
	if err != nil {
		return Range{}, fmt.Errorf("invalid range lower bound: %w", err)
	}
	upper, err := strconv.Atoi(values[1])
	if err != nil {
		return Range{}, fmt.Errorf("invalid range upper bound: %w", err)
	}
	return Range{Lower: lower, Upper: upper}, nil
}

func (r Range) IsFruitFresh(fruit string) bool {
	fInt, _ := strconv.Atoi(fruit)
	return fInt >= r.Lower && fInt <= r.Upper
}

func (r Range) GetWidth() int {
	return r.Upper - r.Lower + 1
}

func (r Range) DebugPrint() {
	fmt.Printf("Range: %d - %d\n", r.Lower, r.Upper)
}

func Day05PartA(input string, verbose bool) (int, error) {
	database := strings.Split(input, "\n\n")

	ranges := strings.Split(database[0], "\n")
	fruits := strings.Split(database[1], "\n")

	var rangeObjs []Range
	for _, r := range ranges {
		newRange, err := NewRange(r)

		if err != nil {
			return 0, err
		}

		rangeObjs = append(rangeObjs, newRange)
	}

	totalValid := 0

	for _, f := range fruits {
		for _, r := range rangeObjs {
			if r.IsFruitFresh(f) {
				if verbose {
					fmt.Printf("Found fresh fruit: %s\n", f)
				}

				totalValid++
				break
			}
		}
	}

	return totalValid, nil
}

func Day05PartB(input string, verbose bool) (int, error) {
	database := strings.Split(input, "\n\n")

	ranges := strings.Split(database[0], "\n")
	var mergedRanges []Range

	var rangeObjs []Range
	for _, r := range ranges {
		newRange, err := NewRange(r)

		if err != nil {
			return 0, err
		}

		rangeObjs = append(rangeObjs, newRange)
	}

	rangeCmp := func(a, b Range) int {
		if a.Lower < b.Lower {
			return -1
		}

		if a.Lower == b.Lower {
			return 0
		}

		return 1
	}

	slices.SortFunc(rangeObjs, rangeCmp)

	mergedRanges = append(mergedRanges, rangeObjs[0])

	prev := mergedRanges[0]
	editingIndex := 0
	for _, r := range rangeObjs[1:] {
		if r.Lower <= prev.Upper {
			mergedRanges[editingIndex].Upper = max(mergedRanges[editingIndex].Upper, r.Upper)
		} else {
			mergedRanges = append(mergedRanges, r)
			editingIndex++
		}

		prev = mergedRanges[editingIndex]
	}

	AccSum := 0

	for _, r := range mergedRanges {
		AccSum += r.GetWidth()

		if verbose {
			r.DebugPrint()
		}
	}

	return AccSum, nil
}
