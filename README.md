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
SelectionSort        5.856µs
BubbleSortFlag       6.448µs
BubbleSort           10.225µ

=== Size: 1000 (repeats: 10) ===
Algorithm            Avg Time
SelectionSort        435.489µs
BubbleSortFlag       727.172µs
BubbleSort           900.598µs

=== Size: 10000 (repeats: 10) ===
Algorithm            Avg Time
SelectionSort        39.262782ms
BubbleSortFlag       84.817661ms
BubbleSort           111.625699ms


```

## Tests

```bash
go test ./tests
```

## Implemented Algorithms

1. **BubbleSort** – Classic Bubble Sort algorithm.
2. **BubbleSortFlag** – Optimized version of Bubble Sort that uses a flag to reduce unnecessary iterations.
3. **SelectionSort** – Finds the smallest element in the unsorted part and swaps it with the first unsorted element.
