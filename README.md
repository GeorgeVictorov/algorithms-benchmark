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
go run .
```

## Output

```

=== Size: 100 (repeats: 10) ===
Algorithm            Avg Time
ShellSort            1.624µs
QuickSortInplace     1.624µs
InsertionSort         2.034µs
SelectionSort        6.684µs
BubbleSortFlag       8.235µs
BubbleSort           13.48µs
QuickSort            39.358µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Avg Time
QuickSortInplace     49.171µs
ShellSort            71.233µs
InsertionSort         259.251µs
QuickSort            450.355µs
SelectionSort        630.498µs
BubbleSortFlag       814.672µs
BubbleSort           1.014293ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Avg Time
QuickSortInplace     782.736µs
ShellSort            1.450137ms
QuickSort            3.628658ms
InsertionSort         14.376545ms
SelectionSort        43.040804ms
BubbleSortFlag       93.21458ms
BubbleSort           115.37261ms

Benchmark completed in: 1.165245154s
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
5. **ShellSort** - An optimized Insertion Sort that compares elements far apart, reducing the number of swaps.
6. **QuickSort** – Recursively sorts by partitioning the array into elements less or greater than a pivot.
7. **QuickSortInplace** – In-place version of QuickSort that sorts the array without creating new slices.
