package services

import (
	"math"
	"strconv"
	"strings"
)

func FloorDiv100(x int) int {
	return int(math.Floor(float64(x) / 100.0))
}

func Fold(nums []int, initial int, fn func(acc, val int) int) int {
	acc := initial
	for _, n := range nums {
		acc = fn(acc, n)
	}
	return acc
}

func DigitsToNum(digits []int) int {
	var sb strings.Builder
	for _, v := range digits {
		sb.WriteByte(byte('0' + v))
	}
	sum, _ := strconv.Atoi(sb.String())
	return sum
}
