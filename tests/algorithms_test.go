package tests

import (
    "reflect"
    "testing"
    "github.com/GeorgeVictorov/algorithms-benchmark/algorithm"
)

func TestBubbleSort(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
        {input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
        {input: []int{5, -1, 3}, expected: []int{-1, 3, 5}},
        {input: []int{}, expected: []int{}},
    }

    for _, c := range cases {
        result := algorithm.BubbleSort(c.input)
        if !reflect.DeepEqual(result, c.expected) {
            t.Errorf("BubbleSort(%v) == %v, expected %v", c.input, result, c.expected)
        }
    }
}

