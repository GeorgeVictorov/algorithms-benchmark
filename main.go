package main

import (
	"github.com/GeorgeVictorov/algorithms-benchmark/algorithm/sorting"
	"github.com/GeorgeVictorov/algorithms-benchmark/util"
)

const Repeats = 10

func main() {
	sorts := []util.Sorter{
		util.SortAlgorithmFunc{N: "BubbleSort", F: sorting.BubbleSort},
		util.SortAlgorithmFunc{N: "BubbleSortFlag", F: sorting.BubbleSortFlag},
		util.SortAlgorithmFunc{N: "SelectionSort", F: sorting.SelectionSort},
		util.SortAlgorithmFunc{N: "InsertionSort", F: sorting.InsertionSort},
		util.SortAlgorithmFunc{N: "ShellSort", F: sorting.ShellSort},
		util.SortAlgorithmFunc{N: "QuickSort", F: sorting.QuickSort},
		util.SortAlgorithmFunc{N: "QuickSortInplace", F: sorting.QuickSortInplace},
		util.SortAlgorithmFunc{N: "MergeSort", F: sorting.MergeSort},
		util.SortAlgorithmFunc{N: "MergeSortParallel", F: sorting.MergeSortParallel},
		util.SortAlgorithmFunc{N: "HeapSort", F: sorting.HeapSort},
	}

	arrSizes := []int{100, 1000, 10000}
	generator := &util.RandomArrayGenerator{}

	util.BenchmarkSorts(sorts, arrSizes, Repeats, generator)
}
