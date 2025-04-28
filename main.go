package main

import (
	"github.com/GeorgeVictorov/algorithms-benchmark/algorithm"
	"github.com/GeorgeVictorov/algorithms-benchmark/util"
)

const Repeats = 10

func main() {
	sorts := []util.Sorter{
		util.SortAlgorithmFunc{N: "BubbleSort", F: algorithm.BubbleSort},
		util.SortAlgorithmFunc{N: "BubbleSortFlag", F: algorithm.BubbleSortFlag},
		util.SortAlgorithmFunc{N: "SelectionSort", F: algorithm.SelectionSort},
		util.SortAlgorithmFunc{N: "InsertionSort", F: algorithm.InsertionSort},
		util.SortAlgorithmFunc{N: "ShellSort", F: algorithm.ShellSort},
		util.SortAlgorithmFunc{N: "QuickSort", F: algorithm.QuickSort},
		util.SortAlgorithmFunc{N: "QuickSortInplace", F: algorithm.QuickSortInplace},
	}

	arrSizes := []int{100, 1000, 10000}
	generator := &util.RandomArrayGenerator{}

	util.BenchmarkSorts(sorts, arrSizes, Repeats, generator)
}
