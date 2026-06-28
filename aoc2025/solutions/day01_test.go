package solutions

import (
	"os"
	"testing"
)

func TestDay01PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/1-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day01PartA(string(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 3
	if got != want {
		t.Errorf("Day01PartA() = %d, want %d", got, want)
	}
}

func TestDay01PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/1-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day01PartB(string(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 6
	if got != want {
		t.Errorf("Day01PartB() = %d, want %d", got, want)
	}
}
