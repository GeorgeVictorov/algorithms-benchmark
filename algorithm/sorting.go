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
