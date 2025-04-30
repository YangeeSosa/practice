package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	line := scanner.Text()
	elements := strings.Fields(line)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(elements[i])
	}

	minDiff := abs(arr[0] - arr[1])
	a, b := arr[0], arr[1]

	for i := 1; i < n-1; i++ {
		currentDiff := abs(arr[i] - arr[i+1])
		if currentDiff < minDiff {
			minDiff = currentDiff
			a, b = arr[i], arr[i+1]
		}
	}

	fmt.Printf("%d %d\n", a, b)
}
