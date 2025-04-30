# Algorithms Benchmark

A project to test the performance of various sorting algorithms.

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
Algorithm            Size       Avg Time
QuickSortInplace     100        1.452µs
ShellSort            100        1.928µs
InsertionSort        100        1.979µs
SelectionSort        100        6.098µs
BubbleSortFlag       100        7.129µs
BubbleSort           100        10.652µs
MergeSort            100        15.679µs
QuickSort            100        47.186µs
MergeSortParallel    100        134.97µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     1000       45.979µs
ShellSort            1000       98.449µs
MergeSort            1000       266.012µs
MergeSortParallel    1000       267.944µs
InsertionSort        1000       292.725µs
QuickSort            1000       550.24µs
SelectionSort        1000       696.185µs
BubbleSortFlag       1000       971.422µs
BubbleSort           1000       1.177088ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     10000      1.529824ms
MergeSortParallel    10000      2.086344ms
ShellSort            10000      2.293575ms
MergeSort            10000      3.409982ms
QuickSort            10000      5.326074ms
InsertionSort        10000      15.526453ms
SelectionSort        10000      45.138135ms
BubbleSortFlag       10000      92.03053ms
BubbleSort           10000      115.262454ms

Benchmark completed in: 1.167011367s
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
8. **MergeSort** – A divide-and-conquer algorithm that recursively splits and merges the array in sorted order.
9. **MergeSortParallel** – A parallelized version of MergeSort that sorts array halves concurrently using goroutines for better performance on multi-core CPUs.
10. **HeapSort** – Sorts by building a binary heap and extracting the max (or min) element, with O(n log n) time complexity.
