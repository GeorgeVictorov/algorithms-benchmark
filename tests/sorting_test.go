package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/GeorgeVictorov/algorithms-benchmark/algorithm/sorting"
)

var testCases = []struct {
	input    []int
	expected []int
}{
	{input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
	{input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	{input: []int{5, -1, 3}, expected: []int{-1, 3, 5}},
	{input: []int{}, expected: []int{}},
	{input: []int{0, 0, 0}, expected: []int{0, 0, 0}},
}

func runSortTest(t *testing.T, name string, sortFunc func([]int) []int) {
	for _, c := range testCases {
		testname := fmt.Sprintf("input=%v", c.input)
		c := c
		t.Run(testname, func(t *testing.T) {
			inputCopy := make([]int, len(c.input))
			copy(inputCopy, c.input)

			result := sortFunc(inputCopy)
			if !reflect.DeepEqual(result, c.expected) {
				t.Errorf("%s(%v) == %v, expected %v", name, c.input, result, c.expected)
			}
		})
	}
}

func TestBubbleSort(t *testing.T)       { runSortTest(t, "BubbleSort", sorting.BubbleSort) }
func TestBubbleSortFlag(t *testing.T)   { runSortTest(t, "BubbleSortFlag", sorting.BubbleSortFlag) }
func TestSelectionSort(t *testing.T)    { runSortTest(t, "SelectionSort", sorting.SelectionSort) }
func TestInsertionSort(t *testing.T)    { runSortTest(t, "InsertionSort", sorting.InsertionSort) }
func TestShellSort(t *testing.T)        { runSortTest(t, "ShellSort", sorting.ShellSort) }
func TestQuickSort(t *testing.T)        { runSortTest(t, "QuickSort", sorting.QuickSort) }
func TestQuickSortInplace(t *testing.T) { runSortTest(t, "QuickSortInplace", sorting.QuickSortInplace) }
func TestMergeSort(t *testing.T)        { runSortTest(t, "MergeSort", sorting.MergeSort) }
func TestMergeSortP(t *testing.T)       { runSortTest(t, "MergeSortParallel", sorting.MergeSortParallel) }
func TestHeapSort(t *testing.T)         { runSortTest(t, "HeapSort", sorting.HeapSort) }
func TestTimSort(t *testing.T)          { runSortTest(t, "TimSort", sorting.TimSort) }

func BenchmarkBubbleSort(b *testing.B) {
	for b.Loop() {
		for _, c := range testCases {
			inputCopy := make([]int, len(c.input))
			copy(inputCopy, c.input)
			sorting.BubbleSort(inputCopy)
		}
	}
}
