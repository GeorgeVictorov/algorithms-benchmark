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

=== Size: 100 (repeats: 10) ===
Algorithm            Avg Time
SelectionSort        6.942µs
BubbleSortFlag       7.986µs
BubbleSort           12.03µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Avg Time
SelectionSort        511.745µs
BubbleSortFlag       802.67µs
BubbleSort           1.063225ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Avg Time
SelectionSort        41.790299ms
BubbleSortFlag       89.05419ms
BubbleSort           111.703024ms

Benchmark completed in: 1.12888931s
```

## Tests

```bash
go test ./tests
```

## Implemented Algorithms

1. **BubbleSort** – Classic Bubble Sort algorithm.
2. **BubbleSortFlag** – Optimized version of Bubble Sort that uses a flag to reduce unnecessary iterations.
3. **SelectionSort** – Finds the smallest element in the unsorted part and swaps it with the first unsorted element.
