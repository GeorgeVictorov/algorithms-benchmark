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
InsertionSort        100        2.119µs
HeapSort             100        2.483µs
QuickSortInplace     100        3.735µs
ShellSort            100        4.14µs
SelectionSort        100        6.557µs
BubbleSortFlag       100        7.29µs
BubbleSort           100        11.526µs
MergeSort            100        15.777µs
QuickSort            100        43.264µs
MergeSortParallel    100        161.278µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     1000       47.305µs
ShellSort            1000       67.29µs
HeapSort             1000       129.218µs
MergeSortParallel    1000       249.818µs
MergeSort            1000       254.639µs
InsertionSort        1000       287.493µs
QuickSort            1000       572.126µs
SelectionSort        1000       751.647µs
BubbleSortFlag       1000       964.797µs
BubbleSort           1000       1.169481ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     10000      1.35887ms
ShellSort            10000      2.163767ms
HeapSort             10000      2.454318ms
MergeSortParallel    10000      2.556763ms
MergeSort            10000      3.588591ms
QuickSort            10000      5.90826ms
InsertionSort        10000      16.209934ms
SelectionSort        10000      50.045062ms
BubbleSortFlag       10000      94.186956ms
BubbleSort           10000      117.202896ms

Benchmark completed in: 1.186495821s
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
