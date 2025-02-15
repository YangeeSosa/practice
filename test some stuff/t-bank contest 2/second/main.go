package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		result := findBouquetCost(a)
		fmt.Println(result)
	}
}

func findBouquetCost(a int64) int64 {
	maxPower := 60
	bestCost := int64(-1)

	for i := 0; i < maxPower; i++ {
		for j := i + 1; j < maxPower; j++ {
			for k := j + 1; k < maxPower; k++ {
				cost := int64((1 << i) + (1 << j) + (1 << k))
				if cost <= a && cost > bestCost {
					bestCost = cost
				}
			}
		}
	}
	return bestCost
}
