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
InsertionSort        100        2.536µs
ShellSort            100        2.544µs
HeapSort             100        3.755µs
QuickSortInplace     100        6.613µs
SelectionSort        100        8.574µs
BubbleSortFlag       100        9.311µs
BubbleSort           100        14.41µs
MergeSort            100        14.896µs
QuickSort            100        50.639µs
MergeSortParallel    100        125.473µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     1000       66.024µs
ShellSort            1000       96.182µs
HeapSort             1000       148.448µs
MergeSort            1000       323.465µs
MergeSortParallel    1000       328.343µs
InsertionSort        1000       384.215µs
QuickSort            1000       597.873µs
SelectionSort        1000       849.277µs
BubbleSortFlag       1000       1.117405ms
BubbleSort           1000       1.412816ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     10000      1.75203ms
ShellSort            10000      2.466562ms
HeapSort             10000      2.668541ms
MergeSortParallel    10000      2.78117ms
MergeSort            10000      3.894474ms
QuickSort            10000      5.968342ms
InsertionSort        10000      16.335814ms
SelectionSort        10000      49.982097ms
BubbleSortFlag       10000      93.305913ms
BubbleSort           10000      117.043312ms

Benchmark completed in: 1.187922743s
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
