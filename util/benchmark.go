package util

import (
	"fmt"
	"sort"
	"time"
)

type SortAlgorithm struct {
	Name string
	Func func([]int) []int
}

type BenchmarkResult struct {
	Name    string
	AvgTime time.Duration
}

func BenchmarkSorts(sorts []SortAlgorithm, sizes []int, repeats int) {
	benchmarkStart := time.Now()

	for _, size := range sizes {
		fmt.Printf("\n=== Size: %d (repeats: %d) ===\n", size, repeats)
		fmt.Printf("%-20s %-15s\n", "Algorithm", "Avg Time")

		arr := RandomArray(size, 10000)

		results := make([]BenchmarkResult, 0, len(sorts))

		for _, sortAlg := range sorts {
			var totalTime time.Duration

			for i := 0; i < repeats; i++ {
				copiedArr := make([]int, len(arr))
				copy(copiedArr, arr)

				start := time.Now()
				sortAlg.Func(copiedArr)
				elapsed := time.Since(start)

				totalTime += elapsed
			}

			avgTime := totalTime / time.Duration(repeats)
			results = append(results, BenchmarkResult{
				Name:    sortAlg.Name,
				AvgTime: avgTime,
			})
		}

		sort.Slice(results, func(i, j int) bool {
			return results[i].AvgTime < results[j].AvgTime
		})

		for _, res := range results {
			fmt.Printf("%-20s %-15v\n", res.Name, res.AvgTime)
		}
	}

	benchmarkDuration := time.Since(benchmarkStart)
	fmt.Printf("\nBenchmark completed in: %v\n", benchmarkDuration)
}
