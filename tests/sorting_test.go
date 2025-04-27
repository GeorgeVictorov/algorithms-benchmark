package tests

import (
	"github.com/GeorgeVictorov/algorithms-benchmark/algorithm"
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

		result := algorithm.BubbleSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("BubbleSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestBubbleSortFlag(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := algorithm.BubbleSortFlag(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("BubbleSortFlag(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := algorithm.SelectionSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("SelectionSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := algorithm.InsertionSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("InsertionSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestShellSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := algorithm.ShellSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("ShellSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := algorithm.QuickSort(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("TestQuickSort(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestQuickSortInplace(t *testing.T) {
	for _, c := range testCases {
		inputCopy := make([]int, len(c.input))
		copy(inputCopy, c.input)

		result := algorithm.QuickSortInplace(inputCopy)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("TestQuickSortInplace(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}
