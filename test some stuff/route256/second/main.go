package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		nStr := scanner.Text()
		n, _ := strconv.ParseUint(nStr, 10, 64)

		if n == 0 {
			fmt.Println(1)
			continue
		}

		sum := 0

		maxZero := 1
		currentPower := uint64(1)
		m := 0

		for {
			nextPower := currentPower * 10

			if nextPower <= n {
				zeros := m + 1
				if zeros > maxZero {
					maxZero = zeros
				}
			}

			if m >= 1 {
				a := 9
				for a >= 1 {
					candidate := uint64(a) * currentPower
					if candidate <= n {
						if m > maxZero {
							maxZero = m
						}
						break
					}
					a--
				}
			}

			if nextPower > n {
				break
			}

			currentPower = nextPower
			m++
		}

		sum += maxZero

		lenN := len(nStr)

		for d := 1; d <= 9; d++ {
			maxK := 0
			for k := lenN; k >= 1; k-- {
				var pow10 uint64 = 1
				for j := 0; j < k; j++ {
					pow10 *= 10
				}
				numerator := pow10 - 1
				num := uint64(d) * numerator / 9
				if num <= n {
					maxK = k
					break
				}
			}
			sum += maxK
		}

		fmt.Println(sum)
	}
}
