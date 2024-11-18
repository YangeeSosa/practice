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
	n, _ := strconv.Atoi(scanner.Text())

	times := make([]int64, n+1)
	dependencies := make([][]int, n+1)
	inDegree := make([]int, n+1)

	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		ti, _ := strconv.ParseInt(line[0], 10, 64)
		times[i] = ti

		for _, dep := range line[1:] {
			depInt, _ := strconv.Atoi(dep)
			dependencies[depInt] = append(dependencies[depInt], i)
			inDegree[i]++
		}
	}

	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range dependencies[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
			if times[neighbor] < times[node]+times[neighbor] {
				times[neighbor] = times[node] + times[neighbor]
			}
		}
	}

	maxTime := int64(0)
	for i := 1; i <= n; i++ {
		if times[i] > maxTime {
			maxTime = times[i]
		}
	}

	fmt.Println(maxTime)
}
