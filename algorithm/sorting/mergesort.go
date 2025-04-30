package sorting

import (
	"math"
	"runtime"
)

func MergeSortParallel(arr []int) []int {
	cpuCount := runtime.NumCPU()
	maxDepth := 2 * int(math.Log2(float64(cpuCount)))
	return mergeSortParallelInternal(arr, maxDepth)
}

func mergeSortParallelInternal(arr []int, depth int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if depth <= 0 {
		mid := len(arr) / 2
		left := mergeSortParallelInternal(arr[:mid], 0)
		right := mergeSortParallelInternal(arr[mid:], 0)
		return merge(left, right)
	}

	mid := len(arr) / 2
	leftChan := make(chan []int)
	rightChan := make(chan []int)

	go func() { leftChan <- mergeSortParallelInternal(arr[:mid], depth-1) }()
	go func() { rightChan <- mergeSortParallelInternal(arr[mid:], depth-1) }()

	left := <-leftChan
	right := <-rightChan
	return merge(left, right)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
