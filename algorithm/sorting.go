package algorithm

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
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
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		n--
	}
	return arr
}

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
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

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		curVal := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > curVal {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curVal
	}
	return arr
}

func ShellSort(arr []int) []int {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}
	return arr
}

func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	pivot := arr[n-1]
	smaller := make([]int, 0)
	greater := make([]int, 0)

	for i := 0; i < n-1; i++ {
		if arr[i] <= pivot {
			smaller = append(smaller, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}

	smaller = QuickSort(smaller)
	greater = QuickSort(greater)

	return append(append(smaller, pivot), greater...)
}

func QuickSortInplace(arr []int) []int {
	return quickSortInplace(arr, 0, len(arr)-1)
}

func quickSortInplace(arr []int, low int, high int) []int {
	if low < high {
		pi := partition(arr, low, high)

		quickSortInplace(arr, low, pi-1)
		quickSortInplace(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func MergeSortChan(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	ch := make(chan int)
	go mergeSortChan(arr, ch)

	sorted := make([]int, 0, len(arr))
	for v := range ch {
		sorted = append(sorted, v)
	}
	return sorted
}

func mergeSortChan(arr []int, ch chan int) {
	defer close(ch)
	if len(arr) < 2 {
		ch <- arr[0]
		return
	}

	left := make(chan int)
	right := make(chan int)

	go mergeSortChan(arr[:len(arr)/2], left)
	go mergeSortChan(arr[len(arr)/2:], right)

	mergech(left, right, ch)
}

func mergech(left, right, c chan int) {
	val, ok := <-left
	val2, ok2 := <-right

	for ok && ok2 {
		if val < val2 {
			c <- val
			val, ok = <-left
		} else {
			c <- val2
			val2, ok2 = <-right
		}
	}
	for ok {
		c <- val
		val, ok = <-left
	}
	for ok2 {
		c <- val2
		val2, ok2 = <-right
	}
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
