# Advent of Code 2025

My Solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Usage

Run a specific day:

```sh
go run . <day>
```

For example:

```sh
go run . 3
```

## Adding a new day

```sh
go run . new <day>
```

This scaffolds `cmd/dayNN.go`, `solutions/dayNN.go`, `solutions/dayNN_test.go`, and empty input files under `inputs/`.

## Running tests

```sh
go test ./solutions/...
```

Place sample input in `inputs/<day>-sample.txt` before running tests.

## Scripts

| Script | Description |
|--------|-------------|
| `scripts/run_tests.sh` | Run the full test suite |
| `scripts/gen_report.sh` | Run tests with coverage and open an HTML coverage report |
