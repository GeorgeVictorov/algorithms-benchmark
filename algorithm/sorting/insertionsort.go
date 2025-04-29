package sorting

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
