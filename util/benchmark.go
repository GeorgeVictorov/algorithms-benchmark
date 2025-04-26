package util

import (
    "fmt"
    "time"
    "github.com/GeorgeVictorov/algorithms-benchmark/util"
)

type SortAlgorithm struct (
    Name string
    Func func([]int) [] int
)

func BechmarkSorts(sorts []SortAlgorithm, sizes []int) {
    fmt.Printf("%-20s %-10s %-10s\n", "Algorithm", "Size", "Time")

    for _, size := range sizes {
        arr := RandomArray(size, 10000)
    
        for _, sort := range sorts {
            copiedArr := make([]int, len(arr))
            copy(copiedArr, arr)

            start := time.Now()
            sort.Func(copiedArr)
            end := time.Since(start)

            fmt.Printf("%-20s %-10d %-10v\n", sort.Name, size, end)
        }
    }
}
