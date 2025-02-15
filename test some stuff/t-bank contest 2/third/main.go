package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])

	scanner.Scan()
	distances := strings.Split(scanner.Text(), " ")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(distances[i])
	}

	changes := 0
	goodDays := 0

	for i := 2; i < n; i++ {
		if a[i] >= a[0] && a[i] <= a[1] {
			goodDays++
		} else {
			if a[i] < a[0] {
				changes += a[0] - a[i]
				a[i] = a[0]
			} else if a[i] > a[1] {
				changes += a[i] - a[1]
				a[i] = a[1]
			}
			goodDays++
		}
	}

	if goodDays >= m {
		fmt.Println(changes)
	} else {
		fmt.Println(-1)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
