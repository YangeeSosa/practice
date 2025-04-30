package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(qualityControl(arr))
}

func qualityControl(arr []int) int {
	result := -1
	for i := range arr {
		if arr[i]%2 == 0 {
			result = arr[i]
		}
	}

	return result
}
