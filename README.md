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
ShellSort            1.696µs
InsetionSort         2.263µs
SelectionSort        6.073µs
BubbleSortFlag       12.181µs
BubbleSort           13.683µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Avg Time
ShellSort            62.018µs
InsetionSort         140.283µs
SelectionSort        541.42µs
BubbleSortFlag       865.033µs
BubbleSort           1.087754ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Avg Time
ShellSort            820.3µs
InsetionSort         12.89484ms
SelectionSort        41.725315ms
BubbleSortFlag       91.638502ms
BubbleSort           112.496346ms


Benchmark completed in: 1.137055023s
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
