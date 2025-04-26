# Algorithms Benchmark

A project to test the performance of various sorting algorithms.

## Structure

- `algorithm/` — sorting algorithms
- `util/` — random data generator
- `tests/` — tests
- `main.go` — timeit

## Run

```bash
go run main.go
```

## Output

```
Algorithm            Size       Time
BubbleSort           100        18.266µs
BubbleSortFlag       100        13.614µs
BubbleSort           1000       1.114119ms
BubbleSortFlag       1000       808.225µs
BubbleSort           10000      113.406789ms
BubbleSortFlag       10000      87.107507ms
```

## Tests

```bash
go test ./tests
```
