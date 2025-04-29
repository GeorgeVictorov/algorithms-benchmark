package tests

import (
	"github.com/GeorgeVictorov/algorithms-benchmark/algorithm/sorting"
	"reflect"
	"testing"
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

func TestBubbleSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.BubbleSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("BubbleSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestBubbleSortFlag(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.BubbleSortFlag(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("BubbleSortFlag(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.SelectionSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("SelectionSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.InsertionSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("InsertionSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestShellSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.ShellSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("ShellSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.QuickSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("TestQuickSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestQuickSortInplace(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.QuickSortInplace(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("TestQuickSortInplace(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestMergeSortParallel(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.MergeSortParallel(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("MergeSortParallel(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := sorting.MergeSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("MergeSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}
