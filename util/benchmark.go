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

type ArrayGenerator interface {
    Generate(size, maxValue int) []int
}

type RandomArrayGenerator struct{}
func (r *RandomArrayGenerator) Generate(size, maxValue int) []int {
    return RandomArray(size, maxValue)
}

func BenchmarkSorts(sorts []SortAlgorithm, sizes []int, repeats int, generator ArrayGenerator) {
    benchmarkStart := time.Now()

    for _, size := range sizes {
        arr := generator.Generate(size, 10000)
        results := benchmarkForSize(sorts, arr, size, repeats)
        printResults(size, repeats, results)
    }

    benchmarkDuration := time.Since(benchmarkStart)
    fmt.Printf("\nBenchmark completed in: %v\n", benchmarkDuration)
}

func benchmarkForSize(sorts []SortAlgorithm, arr []int, size, repeats int) []BenchmarkResult {
    results := make([]BenchmarkResult, len(sorts))
    var wg sync.WaitGroup
    wg.Add(len(sorts))

    for i, sortAlg := range sorts {
        go runBenchmark(i, sortAlg, arr, repeats, results, &wg, size)
    }

    wg.Wait()

    sort.Slice(results, func(i, j int) bool {
        return results[i].AvgTime < results[j].AvgTime
    })

    return results
}

func runBenchmark(i int, sortAlg SortAlgorithm, arr []int, repeats int, results []BenchmarkResult, wg *sync.WaitGroup, size int) {
    defer wg.Done()
    avgTime := benchmarkSort(sortAlg, arr, repeats)
    results[i] = BenchmarkResult{
        Name:    sortAlg.Name,
        Size:    size,
        AvgTime: avgTime,
    }
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

func printResults(size, repeats int, results []BenchmarkResult) {
    fmt.Printf("\n=== Size: %d (repeats: %d) ===\n", size, repeats)
    fmt.Printf("%-20s %-10s %-15s\n", "Algorithm", "Size", "Avg Time")

    for _, res := range results {
        fmt.Printf("%-20s %-10d %-15v\n", res.Name, res.Size, res.AvgTime)
    }
}
