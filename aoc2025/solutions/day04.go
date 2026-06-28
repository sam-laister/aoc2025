package solutions

import (
	"bytes"
	"fmt"
	"strings"
)

func Day04PartA(input string, verbose bool) (int, error) {
	scene := strings.Split(strings.TrimSpace(input), "\n")

	yMax := len(scene) - 1
	xMax := len(scene[0]) - 1

	total := 0

	for y, row := range scene {
		for x, cell := range row {
			if cell == '.' {
				continue
			}

			totalSurrounding := 0

			for yOffset := -1; yOffset <= 1; yOffset++ {
				for xOffset := -1; xOffset <= 1; xOffset++ {
					if yOffset == 0 && xOffset == 0 {
						continue
					}

					yPos := y + yOffset
					xPos := x + xOffset

					if yPos < 0 || yPos > yMax {
						continue
					}

					if xPos < 0 || xPos > xMax {
						continue
					}

					if scene[yPos][xPos] == '@' {
						totalSurrounding += 1
					}
				}
			}

			if verbose {
				fmt.Printf("At x: %d, y: %d, found: %d\n", x, y, totalSurrounding)
			}

			if totalSurrounding < 4 {
				total++
			}
		}
	}

	return total, nil
}

func Day04PartB(input string, verbose bool) (int, error) {
	scene := strings.Split(strings.TrimSpace(input), "\n")

	curr := make([][]byte, len(scene))
	for i, row := range scene {
		curr[i] = []byte(row)
	}

	yMax := len(curr) - 1
	xMax := len(curr[0]) - 1

	total := 0

	iteration := 0

	for {
		if verbose {
			fmt.Printf("Iteration %d\n", iteration)
		}

		next := deepCopyOf(curr)
		changed := false

		for y, row := range curr {
			for x, cell := range row {
				if cell == '.' {
					continue
				}

				totalSurrounding := 0

				for yOffset := -1; yOffset <= 1; yOffset++ {
					for xOffset := -1; xOffset <= 1; xOffset++ {
						if yOffset == 0 && xOffset == 0 {
							continue
						}

						yPos := y + yOffset
						xPos := x + xOffset

						if yPos < 0 || yPos > yMax {
							continue
						}

						if xPos < 0 || xPos > xMax {
							continue
						}

						if curr[yPos][xPos] == '@' {
							totalSurrounding += 1
						}
					}
				}

				if verbose {
					fmt.Printf("At x: %d, y: %d, found: %d\n", x, y, totalSurrounding)
				}

				if totalSurrounding < 4 {
					total++

					next[y][x] = '.'
					changed = true
				}
			}
		}

		curr = next
		iteration++

		if !changed {
			break
		}
	}

	return total, nil
}

func deepCopyOf(grid [][]byte) [][]byte {
	cp := make([][]byte, len(grid))
	for i, row := range grid {
		cp[i] = bytes.Clone(row)
	}
	return cp
}
