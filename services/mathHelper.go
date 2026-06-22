package services

import "math"

func FloorDiv100(x int) int {
	return int(math.Floor(float64(x) / 100.0))
}
