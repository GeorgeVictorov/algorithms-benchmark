package util

import (
	"fmt"
	"time"
)

type SortAlgorithm struct {
	Name string
	Func func([]int) []int
}

func BenchmarkSorts(sorts []SortAlgorithm, sizes []int, repeats int) {
    fmt.Printf("%-20s %-10s %-15s\n", "Algorithm", "Size", "Avg Time")

    for _, size := range sizes {
        arr := RandomArray(size, 10000)

        for _, sort := range sorts {
            var totalTime time.Duration

            for i := 0; i < repeats; i++ {
                copiedArr := make([]int, len(arr))
                copy(copiedArr, arr)

                start := time.Now()
                sort.Func(copiedArr)
                elapsed := time.Since(start)

                totalTime += elapsed
            }

            avgTime := totalTime / time.Duration(repeats)

            fmt.Printf("%-20s %-10d %-15v\n", sort.Name, size, avgTime)
        }
    }
}
