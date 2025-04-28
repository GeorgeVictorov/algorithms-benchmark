package util

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

const MaxValue = 10000

type Sorter interface {
	Name() string
	Sort([]int) []int
}

type SortAlgorithmFunc struct {
	N string
	F func([]int) []int
}

func (s SortAlgorithmFunc) Name() string {
	return s.N
}

func (s SortAlgorithmFunc) Sort(arr []int) []int {
	return s.F(arr)
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

func BenchmarkSorts(sorts []Sorter, sizes []int, repeats int, generator ArrayGenerator) {
	benchmarkStart := time.Now()

	for _, size := range sizes {
		arr := generator.Generate(size, 10000)
		results := benchmarkForSize(sorts, arr, size, repeats)
		printResults(size, repeats, results)
	}

	benchmarkDuration := time.Since(benchmarkStart)
	fmt.Printf("\nBenchmark completed in: %v\n", benchmarkDuration)
}

func benchmarkForSize(sorts []Sorter, arr []int, size, repeats int) []BenchmarkResult {
	var wg sync.WaitGroup
	resultsCh := make(chan BenchmarkResult, len(sorts))

	wg.Add(len(sorts))
	for _, sortAlg := range sorts {
		go runBenchmark(sortAlg, arr, repeats, size, resultsCh, &wg)
	}

	wg.Wait()
	close(resultsCh)

	var results []BenchmarkResult
	for res := range resultsCh {
		results = append(results, res)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].AvgTime < results[j].AvgTime
	})

	return results
}

func runBenchmark(sortAlg Sorter, arr []int, repeats, size int, resultsCh chan<- BenchmarkResult, wg *sync.WaitGroup) {
	defer wg.Done()

	avgTime := benchmarkSort(sortAlg, arr, repeats)
	resultsCh <- BenchmarkResult{
		Name:    sortAlg.Name(),
		Size:    size,
		AvgTime: avgTime,
	}
}

func benchmarkSort(sortAlg Sorter, arr []int, repeats int) time.Duration {
	var totalTime time.Duration

	for i := 0; i < repeats; i++ {
		copiedArr := make([]int, len(arr))
		copy(copiedArr, arr)

		start := time.Now()
		sortAlg.Sort(copiedArr)
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
