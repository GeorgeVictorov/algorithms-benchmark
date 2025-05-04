package sorting

const minRun = 32

func TimSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	for start := 0; start < n; start += minRun {
		end := start + minRun
		if end > n {
			end = n
		}
		insertionSort(arr, start, end)
	}

	for size := minRun; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := min(left+size, n)
			right := min(left+2*size, n)
			mergeT(arr, left, mid, right)
		}
	}

	return arr
}

func insertionSort(arr []int, left, right int) {
	for i := left + 1; i < right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func mergeT(arr []int, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid

	leftPart := arr[left:mid]
	rightPart := arr[mid:right]

	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if leftPart[i] <= rightPart[j] {
			arr[k] = leftPart[i]
			i++
		} else {
			arr[k] = rightPart[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = leftPart[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = rightPart[j]
		j++
		k++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
