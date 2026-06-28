package solutions

import (
	"github.com/sam-laister/aoc2025/services"
	"strconv"
	"strings"
)

func Day01PartA(input string) (int, error) {
	data := strings.Split(input, "\n")

	accTotal := 50
	hitTotal := 0

	for _, i := range data {
		if len(i) == 0 {
			continue
		}

		value, err := strconv.Atoi(i[1:])
		if err != nil {
			continue
		}

		if i[0] == 'L' {
			accTotal -= value
		} else {
			accTotal += value
		}

		if accTotal%100 == 0 {
			hitTotal += 1
		}
	}

	return hitTotal, nil
}

func Day01PartB(input string) (int, error) {
	data := strings.Split(input, "\n")

	accTotal := 50
	hitTotal := 0

	for _, i := range data {
		if len(i) == 0 {
			continue
		}

		start := accTotal

		value, err := strconv.Atoi(i[1:])
		if err != nil {
			continue
		}

		if i[0] == 'L' {
			accTotal -= value
		} else {
			accTotal += value
		}

		end := accTotal

		var hits int
		if end > start {
			hits = services.FloorDiv100(end) - services.FloorDiv100(start)
		} else {
			hits = services.FloorDiv100(start-1) - services.FloorDiv100(end-1)
		}

		hitTotal += hits

	}

	return hitTotal, nil
}
