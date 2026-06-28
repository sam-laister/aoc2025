package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sam-laister/aoc2025/services"
)

type Column struct {
	Values []int
	Symbol string
}

func (c Column) Calculate() (int, error) {
	var mapFunc func(acc, val int) int
	var startValue int

	switch c.Symbol {
	case "*":
		mapFunc = func(acc, val int) int {
			return acc * val
		}
		startValue = 1
	case "+":
		mapFunc = func(acc, val int) int {
			return acc + val
		}
		startValue = 0
	default:
		return 0, fmt.Errorf("Incorrect symbol: %s\n", c.Symbol)
	}

	return services.Fold(c.Values, startValue, mapFunc), nil
}

func (c Column) DebugPrint() {
	calc, _ := c.Calculate()
	fmt.Printf("%+v - %d\n", c, calc)
}

func Day06PartA(input string, verbose bool) (int, error) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	fields := make([]Column, len(strings.Fields(rows[0])))

	for i, row := range rows {
		f := strings.Fields(row)

		if i == len(rows)-1 {
			// final row
			for j, c := range f {
				fields[j].Symbol = c
			}
		} else {
			// regular row
			for j, c := range f {
				intC, _ := strconv.Atoi(c)
				fields[j].Values = append(fields[j].Values, intC)
			}
		}
	}

	total := 0
	for _, f := range fields {
		if verbose {
			f.DebugPrint()
		}

		sum, err := f.Calculate()
		if err != nil {
			return 0, err
		}
		total += sum
	}

	return total, nil
}

func Day06PartB(input string, verbose bool) (int, error) {
	rows := strings.Split(strings.TrimSpace(input), "\n")

	numRows := len(rows) - 1
	width := len(rows[0])

	sep := make([]bool, width)
	for col := range width {
		sep[col] = true
		for row := range numRows {
			if rows[row][col] != ' ' {
				sep[col] = false
				break
			}
		}
	}

	var problems [][]int
	var curr []int
	for col := range width {
		if sep[col] {
			if len(curr) > 0 {
				problems = append(problems, curr)
				curr = nil
			}
		} else {
			curr = append(curr, col)
		}
	}

	if len(curr) > 0 {
		problems = append(problems, curr)
	}

	total := 0
	opRow := rows[len(rows)-1]
	for _, cols := range problems {
		var values []int
		for i := len(cols) - 1; i >= 0; i-- {
			col := cols[i]
			var sb strings.Builder

			for row := range numRows {
				if rows[row][col] != ' ' {
					sb.WriteByte(rows[row][col])
				}
			}

			if sb.Len() > 0 {
				val, _ := strconv.Atoi(sb.String())
				values = append(values, val)
			}
		}

		sum, err := Column{Values: values, Symbol: string(opRow[cols[0]])}.Calculate()
		if err != nil {
			return 0, err
		}

		total += sum
	}

	return total, nil

}
