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
InsetionSort         1.81µs
SelectionSort        5.954µs
BubbleSortFlag       6.385µs
BubbleSort           12.286µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Avg Time
InsetionSort         136.121µs
SelectionSort        438.998µs
BubbleSortFlag       686.381µs
BubbleSort           889.441µ
=== Size: 10000 (repeats: 10) ===
Algorithm            Avg Time
InsetionSort         12.760935ms
SelectionSort        41.364914ms
BubbleSortFlag       86.736508ms
BubbleSort           108.776995ms

Benchmark completed in: 1.098451979s
```

## Tests

```bash
go test ./tests
```

## Implemented Algorithms

1. **BubbleSort** – Classic Bubble Sort algorithm.
2. **BubbleSortFlag** – Optimized version of Bubble Sort that uses a flag to reduce unnecessary iterations.
3. **SelectionSort** – Finds the smallest element in the unsorted part and swaps it with the first unsorted element.
4. **InsertionSort** – Sorts by repeatedly picking the next element and inserting it into its correct position in the array.
