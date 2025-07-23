package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	arr = mergeSort(arr)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]

	return merge(mergeSort(left), mergeSort(right))
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
