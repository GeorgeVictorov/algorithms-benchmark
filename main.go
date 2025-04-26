package main

import (
    "github.com/GeorgeVictorov/algorithms-benchmark/algorithm"
    "github.com/GeorgeVictorov/algorithms-benchmark/util"
)

func main() {
    sorts := []util.SortAlgorithm{
		{Name: "BubbleSort", Func: algorithm.BubbleSort},
		{Name: "BubbleSortFlag", Func: algorithm.BubbleSortFlag},
    }

    arrSizes := []int{100, 1000, 10000}

    util.BenchmarkSorts(sorts, arrSizes, 10)
}
