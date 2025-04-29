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
InsertionSort        100        2.213µs
ShellSort            100        3.135µs
BubbleSort           100        10.12µs
BubbleSortFlag       100        10.589µs
QuickSortInplace     100        10.972µs
SelectionSort        100        11.884µs
MergeSort            100        19.958µs
QuickSort            100        51.024µs
MergeSortParallel    100        162.644µs

=== Size: 1000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     1000       50.663µs
ShellSort            1000       70.655µs
InsertionSort        1000       222.576µs
MergeSortParallel    1000       253.301µs
MergeSort            1000       256.371µs
QuickSort            1000       506.507µs
SelectionSort        1000       633.086µs
BubbleSortFlag       1000       853.834µs
BubbleSort           1000       1.093497ms

=== Size: 10000 (repeats: 10) ===
Algorithm            Size       Avg Time
QuickSortInplace     10000      1.454827ms
MergeSortParallel    10000      2.034581ms
ShellSort            10000      2.040737ms
MergeSort            10000      3.386363ms
QuickSort            10000      5.492465ms
InsertionSort        10000      16.180876ms
SelectionSort        10000      45.624726ms
BubbleSortFlag       10000      89.844705ms
BubbleSort           10000      112.249365ms
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
