package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		if a[i] == -1 {
			if i == 0 {
				a[i] = 1
			} else {
				a[i] = a[i-1] + 1
			}
		}
	}

	for i := 1; i < n; i++ {
		if a[i] <= a[i-1] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}
