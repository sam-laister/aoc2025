package solutions

import (
	"os"
	"testing"
)

func TestDay05PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/5-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day05PartA(string(input), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 3
	if got != want {
		t.Errorf("Day05PartA() = %d, want %d", got, want)
	}
}

func TestDay05PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/5-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day05PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 14
	if got != want {
		t.Errorf("Day05PartB() = %d, want %d", got, want)
	}
}
