package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	var s int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &s)

	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	prefixSum := make([]int64, n+1)
	prefixSum[0] = 0
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + a[i]
	}

	totalSubarrays := int64(n) * int64(n+1) / 2
	sumCuts := int64(0)

	for i := 1; i <= n-1; i++ {
		A := prefixSum[i] - s
		B := prefixSum[i+1] - s

		left := sort.Search(i, func(k int) bool {
			return prefixSum[k] >= A
		})

		right := sort.Search(i, func(k int) bool {
			return prefixSum[k] >= B
		})

		count := right - left
		if count < 0 {
			count = 0
		}

		sumCuts += int64(count) * int64(n-i)
	}

	total := totalSubarrays + sumCuts
	fmt.Println(total)
}
