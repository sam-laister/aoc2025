package solutions

import (
	"os"
	"testing"
)

func TestDay02PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/2-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day02PartA(string(input), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 1227775554
	if got != want {
		t.Errorf("Day01PartA() = %d, want %d", got, want)
	}
}
