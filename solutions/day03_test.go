package solutions

import (
	"os"
	"testing"
)

func TestDay03PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/3-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day03PartA(string(input), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 357
	if got != want {
		t.Errorf("Day03PartA() = %d, want %d", got, want)
	}
}

func TestDay03PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/3-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day03PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 3121910778619
	if got != want {
		t.Errorf("Day03PartB() = %d, want %d", got, want)
	}
}
