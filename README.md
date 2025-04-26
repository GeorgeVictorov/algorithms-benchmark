# Algorithms Benchmark

A project to test the performance of various sorting algorithms.

## Structure

- `algorithm/`
- `util/`
- `tests/`
- `main.go`

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
