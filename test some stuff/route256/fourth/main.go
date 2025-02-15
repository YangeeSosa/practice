package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	t := atoi(scanner.Text())

	for testCase := 0; testCase < t; testCase++ {
		scanner.Scan()
		n := atoi(scanner.Text())

		evenCount := make(map[string]int)
		oddCount := make(map[string]int)
		bothCount := make(map[string]int)

		for i := 0; i < n; i++ {
			scanner.Scan()
			s := scanner.Text()

			var evenKey []byte
			for j := 0; j < len(s); j += 2 {
				evenKey = append(evenKey, s[j])
			}
			ek := string(evenKey)

			var ok string
			if len(s) >= 2 {
				var oddKey []byte
				for j := 1; j < len(s); j += 2 {
					oddKey = append(oddKey, s[j])
				}
				ok = string(oddKey)
			} else {
				ok = "NO_ODD_KEY"
			}

			evenCount[ek]++
			if ok != "NO_ODD_KEY" {
				oddCount[ok]++
				bothCount[ek+"|"+ok]++
			}
		}

		total := calcPairs(evenCount) + calcPairs(oddCount) - calcPairs(bothCount)
		fmt.Println(total)
	}
}

func calcPairs(m map[string]int) int {
	pairs := 0
	for _, cnt := range m {
		pairs += cnt * (cnt - 1) / 2
	}
	return pairs
}

func atoi(s string) int {
	res := 0
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}
