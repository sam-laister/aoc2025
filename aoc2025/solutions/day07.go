package solutions

import (
	"fmt"
	"strings"
)

func Day07PartA(input string, verbose bool) (int, error) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	buffA := make([]string, len(rows))

	copy(buffA, rows)

	rowLen := len(buffA[0])
	rowCount := len(buffA)

	splitTotal := 0

	for rowIndex := range rowCount {
		if rowIndex == 0 {
			continue
		}

		prevRow := buffA[rowIndex-1]
		currRow := strings.Split(buffA[rowIndex], "")

		for col, sym := range prevRow {
			currSym := string(sym)
			if currSym == "|" || currSym == "S" {
				switch string(currRow[col]) {
				case ".":
					currRow[col] = "|"
				case "^":
					splitTotal++
					if col-1 >= 0 {
						currRow[col-1] = "|"
					}
					if col+1 < rowLen {
						currRow[col+1] = "|"
					}
				}
			}
		}

		buffA[rowIndex] = strings.Join(currRow, "")
	}

	if verbose {
		fmt.Printf("%+v\n", strings.Join(buffA, "\n"))
	}

	return splitTotal, nil
}

func Day07PartB(input string, verbose bool) (int, error) {
	rows := strings.Split(strings.TrimSpace(input), "\n")

	rowLen := len(rows[0])
	rowCount := len(rows)

	sCol := strings.Index(rows[0], "S")
	counts := make([]int, rowLen)
	counts[sCol] = 1

	for rowIndex := range rowCount {
		if rowIndex == 0 {
			continue
		}

		next := make([]int, rowLen)
		for col, n := range counts {
			if n == 0 {
				continue
			}

			switch string(rows[rowIndex][col]) {
			case ".":
				next[col] += n
			case "^":
				if col-1 >= 0 {
					next[col-1] += n
				}
				if col+1 < rowLen {
					next[col+1] += n
				}
			}
		}
		counts = next
	}

	total := 0
	for _, n := range counts {
		total += n
	}
	return total, nil
}
