package algorithm

func BubbleSort(arr []int) []int {
    n := len(arr)
    for i:= 0; i < n -1; i++ {
        for j :=0; j < n - 1; j++ {
            if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
            }
        }
    }
    return arr
}

func BubbleSortFlag(arr []int) []int {
    n := len(arr)
    swapped := true

    for swapped {
        swapped = false
        for i:= 0; i < n - 1; i++ {
            if arr[i] > arr[i + 1] {
                arr[i], arr[i + 1] = arr[i + 1], arr[i]
                swapped = true
            }
        }
        n--
    }
    return arr
}

func SelectionSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n - 1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
    return arr
}
        
