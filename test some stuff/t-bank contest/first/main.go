package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findRange(s string, result *[]int) {
	if strings.Contains(s, "-") {
		splitted := strings.Split(s, "-")
		start, _ := strconv.Atoi(splitted[0])
		end, _ := strconv.Atoi(splitted[1])

		for i := start; i <= end; i++ {
			*result = append(*result, i)
		}
	} else {
		num, _ := strconv.Atoi(s)
		*result = append(*result, num)
	}
}

func main() {
	var input string
	fmt.Scan(&input)

	parts := strings.Split(input, ",")

	var result []int

	for _, part := range parts {
		findRange(part, &result)
	}

	sort.Ints(result)

	for i, num := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}

}
