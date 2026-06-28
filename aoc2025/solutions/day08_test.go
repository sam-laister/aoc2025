package solutions

import (
	"os"
	"testing"
)

func TestDay08PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/8-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day08PartA(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0
	if got != want {
		t.Errorf("Day08PartA() = %d, want %d", got, want)
	}
}

func TestDay08PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/8-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day08PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0
	if got != want {
		t.Errorf("Day08PartB() = %d, want %d", got, want)
	}
}
