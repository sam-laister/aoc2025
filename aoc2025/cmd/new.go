package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new <day>",
	Short: "Generate skeleton files for a new day",
	Args:  cobra.ExactArgs(1),
	RunE:  runNew,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func runNew(_ *cobra.Command, args []string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil || n < 1 {
		return fmt.Errorf("day must be a positive integer, got %q", args[0])
	}

	day := args[0]
	padded := fmt.Sprintf("%02d", n)

	r := strings.NewReplacer(
		"{{DAY}}", day,
		"{{PADDED}}", padded,
		"{{VAR}}", "day"+day,
		"{{FUNC}}", "Day"+padded,
	)

	files := []struct {
		path    string
		content string
	}{
		{fmt.Sprintf("cmd/day%s.go", padded), r.Replace(newCmdTmpl)},
		{fmt.Sprintf("solutions/day%s.go", padded), r.Replace(newSolutionTmpl)},
		{fmt.Sprintf("solutions/day%s_test.go", padded), r.Replace(newTestTmpl)},
		{fmt.Sprintf("inputs/%s.txt", day), ""},
		{fmt.Sprintf("inputs/%s-sample.txt", day), ""},
	}

	for _, f := range files {
		if _, err := os.Stat(f.path); err == nil {
			return fmt.Errorf("already exists: %s", f.path)
		}
	}

	for _, f := range files {
		if err := os.WriteFile(f.path, []byte(f.content), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", f.path, err)
		}
		fmt.Println("created", f.path)
	}

	return nil
}

const newCmdTmpl = `package cmd

import (
	"fmt"

	"github.com/sam-laister/aoc2025/services"
	"github.com/sam-laister/aoc2025/solutions"
	"github.com/spf13/cobra"
)

var {{VAR}} = &cobra.Command{
	Use:   "{{DAY}}",
	Short: "Day {{DAY}}",
	RunE: func(cmd *cobra.Command, args []string) error {
		txt, err := services.ReadDay("{{DAY}}")
		if err != nil {
			return err
		}

		partA, err := solutions.{{FUNC}}PartA(txt, false)
		if err != nil {
			return err
		}

		partB, err := solutions.{{FUNC}}PartB(txt, false)
		if err != nil {
			return err
		}

		fmt.Println("Part A:", partA)
		fmt.Println("Part B:", partB)

		return nil
	},
}

func init() {
	rootCmd.AddCommand({{VAR}})
}
`

const newSolutionTmpl = `package solutions

import (
	"strings"
)

func {{FUNC}}PartA(input string, verbose bool) (int, error) {
	_ = strings.Split(input, "\n")
	return 0, nil
}

func {{FUNC}}PartB(input string, verbose bool) (int, error) {
	_ = strings.Split(input, "\n")
	return 0, nil
}
`

const newTestTmpl = `package solutions

import (
	"os"
	"testing"
)

func Test{{FUNC}}PartA(t *testing.T) {
	input, err := os.ReadFile("../inputs/{{DAY}}-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := {{FUNC}}PartA(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0
	if got != want {
		t.Errorf("{{FUNC}}PartA() = %d, want %d", got, want)
	}
}

func Test{{FUNC}}PartB(t *testing.T) {
	input, err := os.ReadFile("../inputs/{{DAY}}-sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, err := {{FUNC}}PartB(string(input), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0
	if got != want {
		t.Errorf("{{FUNC}}PartB() = %d, want %d", got, want)
	}
}
`
