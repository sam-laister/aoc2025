package solutions

import (
	"os"
	"testing"
)

func TestDay07PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/7-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day07PartA(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 21
	if got != want {
		t.Errorf("Day07PartA() = %d, want %d", got, want)
	}
}

func TestDay07PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/7-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := Day07PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 40
	if got != want {
		t.Errorf("Day07PartB() = %d, want %d", got, want)
	}
}
