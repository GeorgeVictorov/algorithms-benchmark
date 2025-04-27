package util

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type SortAlgorithm struct {
	Name string
	Func func([]int) []int
}

type BenchmarkResult struct {
	Name    string
	Size    int
	AvgTime time.Duration
}

func BenchmarkSorts(sorts []SortAlgorithm, sizes []int, repeats int) {
	benchmarkStart := time.Now()

	for _, size := range sizes {
		benchmarkForSize(sorts, size, repeats)
	}

	benchmarkDuration := time.Since(benchmarkStart)
	fmt.Printf("\nBenchmark completed in: %v\n", benchmarkDuration)
}

func benchmarkForSize(sorts []SortAlgorithm, size, repeats int) {
	fmt.Printf("\n=== Size: %d (repeats: %d) ===\n", size, repeats)
	fmt.Printf("%-20s %-15s\n", "Algorithm", "Avg Time")

	arr := RandomArray(size, 10000)

	results := make([]BenchmarkResult, len(sorts))

	var wg sync.WaitGroup
	wg.Add(len(sorts))

	for i, sort := range sorts {
		go func(i int, sortAlg SortAlgorithm) {
			defer wg.Done()
			avgTime := benchmarkSort(sortAlg, arr, repeats)
			results[i] = BenchmarkResult{
				Name:    sortAlg.Name,
				AvgTime: avgTime,
			}
		}(i, sort)
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		return results[i].AvgTime < results[j].AvgTime
	})

	printResults(results)
}

func benchmarkSort(sortAlg SortAlgorithm, arr []int, repeats int) time.Duration {
	var totalTime time.Duration

	for i := 0; i < repeats; i++ {
		copiedArr := make([]int, len(arr))
		copy(copiedArr, arr)

		start := time.Now()
		sortAlg.Func(copiedArr)
		totalTime += time.Since(start)
	}

	return totalTime / time.Duration(repeats)
}

func printResults(results []BenchmarkResult) {
	for _, res := range results {
		fmt.Printf("%-20s %-15v\n", res.Name, res.AvgTime)
	}
}
