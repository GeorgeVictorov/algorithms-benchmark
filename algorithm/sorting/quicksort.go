package sorting

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
