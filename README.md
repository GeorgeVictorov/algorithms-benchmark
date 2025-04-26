# Algorithms Benchmark

A project to test the performance of various sorting algorithms.

## Structure

- `algorithm/`
- `util/`
- `tests/`
- `main.go`

## Installation

```bash
git clone https://github.com/GeorgeVictorov/algorithms-benchmark.git
cd algorithms-benchmark
```

## Run

```bash
go run main.go
```

## Output

```
Algorithm            Size       Avg Time
BubbleSort           100        14.79µs
BubbleSortFlag       100        6.964µs
BubbleSort           1000       865.805µs
BubbleSortFlag       1000       646.54µs
BubbleSort           10000      107.027192ms
BubbleSortFlag       10000      88.49015ms
```

## Tests

```bash
go test ./tests
```

## Implemented Algorithms

1. BubbleSort – Classic Bubble Sort algorithm.
2. BubbleSortFlag – Optimized version of Bubble Sort that uses a flag to reduce unnecessary iterations.
