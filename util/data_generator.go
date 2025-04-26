package utils

import (
    "math/rand"
    "time"
)

func RandomArray(size, maxn int) []int {
    if size <= 0 {
        return []int{}
    }    

    rand.Seed(time.Now().UnixNano())

    arr := make([]int, size)
    for i := range arr {
        arr[i] = rand.Intn(maxn)
    }
    return arr
}
