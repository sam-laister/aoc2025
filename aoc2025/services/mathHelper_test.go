package services

import "testing"

func TestFloorDiv100(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{99, 0},
		{100, 1},
		{101, 1},
		{200, 2},
		{-1, -1},
		{-100, -1},
		{-101, -2},
	}
	for _, tt := range tests {
		got := FloorDiv100(tt.input)
		if got != tt.want {
			t.Errorf("FloorDiv100(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestFold(t *testing.T) {
	sum := func(acc, val int) int { return acc + val }
	mul := func(acc, val int) int { return acc * val }

	tests := []struct {
		name    string
		nums    []int
		initial int
		fn      func(int, int) int
		want    int
	}{
		{"sum empty", []int{}, 0, sum, 0},
		{"sum single", []int{5}, 0, sum, 5},
		{"sum multiple", []int{1, 2, 3, 4}, 0, sum, 10},
		{"sum with initial", []int{1, 2, 3}, 10, sum, 16},
		{"mul empty", []int{}, 1, mul, 1},
		{"mul single", []int{5}, 1, mul, 5},
		{"mul multiple", []int{2, 3, 4}, 1, mul, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fold(tt.nums, tt.initial, tt.fn)
			if got != tt.want {
				t.Errorf("Fold() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestDigitsToNum(t *testing.T) {
	tests := []struct {
		digits []int
		want   int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 123},
		{[]int{0, 0, 1}, 1},
		{[]int{9, 9, 9}, 999},
		{[]int{3, 5, 6}, 356},
	}
	for _, tt := range tests {
		got := DigitsToNum(tt.digits)
		if got != tt.want {
			t.Errorf("DigitsToNum(%v) = %d, want %d", tt.digits, got, tt.want)
		}
	}
}
