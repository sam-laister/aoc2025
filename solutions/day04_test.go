package solutions

import (
	"os"
	"testing"
)

func TestDay04PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/4-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day04PartA(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0
	if got != want {
		t.Errorf("Day04PartA() = %d, want %d", got, want)
	}
}

func TestDay04PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/4-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day04PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0
	if got != want {
		t.Errorf("Day04PartB() = %d, want %d", got, want)
	}
}
