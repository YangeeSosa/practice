package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countDivisors(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	result := 0
	for num := l; num <= r; num++ {
		if num > 1 && !isPrime(num) {
			divisorCount := countDivisors(num)
			if isPrime(divisorCount) {
				result++
			}
		}
	}

	fmt.Println(result)
}
