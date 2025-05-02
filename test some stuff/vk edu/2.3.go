package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var num int
	fmt.Scan(&num)

	fmt.Println(expSearch(arr, num))

}

func expSearch(arr []int, target int) (int, int) {
	border := 1
	lastEl := len(arr) - 1

	for border < lastEl && arr[border] < target {
		if arr[border] == target {
			return border, border * 2
		}
		border *= 2
		if border > lastEl {
			return border / 2, lastEl
		}
	}

	return border / 2, border

}
