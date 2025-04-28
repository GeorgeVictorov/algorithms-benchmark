package main

import (
	"github.com/GeorgeVictorov/algorithms-benchmark/algorithm"
	"github.com/GeorgeVictorov/algorithms-benchmark/util"
)

func main() {
	sorts := []util.SortAlgorithm{
		{Name: "BubbleSort", Func: algorithm.BubbleSort},
		{Name: "BubbleSortFlag", Func: algorithm.BubbleSortFlag},
		{Name: "SelectionSort", Func: algorithm.SelectionSort},
        {Name: "InsertionSort", Func: algorithm.InsertionSort},
        {Name: "ShellSort", Func: algorithm.ShellSort},
        {Name: "QuickSort", Func: algorithm.QuickSort},
        {Name: "QuickSortInplace", Func: algorithm.QuickSortInplace},
	}

	arrSizes := []int{100, 1000, 10000}
    
    generator := &util.RandomArrayGenerator{}

	util.BenchmarkSorts(sorts, arrSizes, 10, generator)
}
