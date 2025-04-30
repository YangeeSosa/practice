package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var target int
	fmt.Scan(&target)

	fmt.Println(binarySearchAndInsert(arr, target))

}

func binarySearchAndInsert(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
