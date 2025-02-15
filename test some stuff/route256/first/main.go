package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Scan(&n, &m)
		rows := (n + 1) / 2
		cols := (m + 1) / 2
		total := rows * cols

		fmt.Println(total)

		for i := 1; i <= n; i += 2 {
			for j := 1; j <= m; j += 2 {
				direction := "R"
				if j == m && m%2 == 0 { // Если столбец чётный и достигнут край
					direction = "L"
				}
				if i%2 == 1 && j%2 == 1 {
					if i == n && j < m {
						direction = "R"
					} else if j == m && i < n {
						direction = "D"
					}
				}
				fmt.Printf("%d %d %s\n", i, j, direction)
			}
		}
	}
}
