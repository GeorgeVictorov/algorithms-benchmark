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
TimSort              100        1.793µs
ShellSort            100        2.393µs
InsertionSort        100        2.907µs
HeapSort             100        3.248µs
QuickSortInplace     100        5.811µs
SelectionSort        100        8.044µs
BubbleSortFlag       100        9.705µs
MergeSort            100        12.629µs
BubbleSort           100        14.252µs
QuickSort            100        54.454µs
MergeSortParallel    100        129.626µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Size       Avg Time
TimSort              1000       24.984µs
QuickSortInplace     1000       66.643µs
ShellSort            1000       87.655µs
HeapSort             1000       222.897µs
MergeSort            1000       273.013µs
MergeSortParallel    1000       313.705µs
InsertionSort        1000       321.808µs
QuickSort            1000       571.443µs
SelectionSort        1000       787.306µs
BubbleSortFlag       1000       983.771µs
BubbleSort           1000       1.259172ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Size       Avg Time
TimSort              10000      733.859µs
QuickSortInplace     10000      1.412321ms
ShellSort            10000      2.108549ms
MergeSortParallel    10000      2.401364ms
HeapSort             10000      2.552457ms
MergeSort            10000      3.541477ms
QuickSort            10000      5.627927ms
InsertionSort        10000      15.930532ms
SelectionSort        10000      49.932765ms
BubbleSortFlag       10000      92.517422ms
BubbleSort           10000      114.506208ms

Benchmark completed in: 1.161568793s
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
11. **TimSort** – A hybrid sorting algorithm combining Insertion Sort and Merge Sort.
