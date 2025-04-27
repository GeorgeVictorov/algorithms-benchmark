package tests

import (
	"github.com/GeorgeVictorov/algorithms-benchmark/util"
	"testing"
)

func TestRandomArraySize(t *testing.T) {
	size := 10
	maxn := 100
	arr := util.RandomArray(size, maxn)
	if len(arr) != size {
		t.Errorf("Expected array size %d, got %d", size, len(arr))
	}
}

func TestRandomArrayEmpty(t *testing.T) {
	size := 0
	maxn := 100
	arr := util.RandomArray(size, maxn)
	if len(arr) != 0 {
		t.Errorf("Expected empty array, got size %d", len(arr))
	}
}

func TestRandomArray_MaxValueRange(t *testing.T) {
	size := 100
	maxn := 100
	arr := util.RandomArray(size, maxn)
	for _, v := range arr {
		if v < 0 || v >= maxn {
			t.Errorf("Value out of range: %d, expected between 0 and %d", v, maxn-1)
		}
	}
}
