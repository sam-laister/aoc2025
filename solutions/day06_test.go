package solutions

import (
	"os"
	"testing"
)

func TestDay06PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/6-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day06PartA(string(input), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 4277556
	if got != want {
		t.Errorf("Day06PartA() = %d, want %d", got, want)
	}
}

func TestDay06PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/6-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day06PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 3263827
	if got != want {
		t.Errorf("Day06PartB() = %d, want %d", got, want)
	}
}
