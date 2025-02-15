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
	scanner.Buffer(make([]byte, 3*1024*1024), 3*1024*1024)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for testCase := 0; testCase < t; testCase++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		nameToPrice := make(map[string]int, n)
		priceCount := make(map[int]int, n)
		priceToName := make(map[int]string, n)

		valid := true
		for i := 0; i < n; i++ {
			scanner.Scan()
			line := scanner.Text()
			idx := strings.IndexByte(line, ' ')
			if idx == -1 || idx == 0 || idx == len(line)-1 {
				valid = false
				for j := i + 1; j < n; j++ {
					scanner.Scan()
				}
				break
			}

			name := line[:idx]
			priceStr := line[idx+1:]

			if !isValidNameFast(name) {
				valid = false
				for j := i + 1; j < n; j++ {
					scanner.Scan()
				}
				break
			}

			if len(priceStr) == 0 || (len(priceStr) > 1 && priceStr[0] == '0') {
				valid = false
				for j := i + 1; j < n; j++ {
					scanner.Scan()
				}
				break
			}

			price, err := strconv.Atoi(priceStr)
			if err != nil || price < 1 || price > 1e9 {
				valid = false
				for j := i + 1; j < n; j++ {
					scanner.Scan()
				}
				break
			}

			nameToPrice[name] = price
			priceCount[price]++
			if priceCount[price] == 1 {
				priceToName[price] = name
			} else {
				delete(priceToName, price)
			}
		}

		if !valid {
			scanner.Scan()
			fmt.Println("NO")
			continue
		}

		scanner.Scan()
		s := scanner.Text()

		parsedPrices := make(map[int]string, len(priceCount))
		parts := strings.Split(s, ",")
		valid = len(parts) > 0

		for _, part := range parts {
			split := strings.Split(part, ":")
			if len(split) != 2 {
				valid = false
				break
			}

			name, priceStr := split[0], split[1]

			if !isValidNameFast(name) || !isValidPriceStrFast(priceStr) {
				valid = false
				break
			}

			price, err := strconv.Atoi(priceStr)
			if err != nil || price < 1 || price > 1e9 {
				valid = false
				break
			}

			expectedPrice, exists := nameToPrice[name]
			if !exists || expectedPrice != price {
				valid = false
				break
			}

			if _, exists := parsedPrices[price]; exists {
				valid = false
				break
			}
			parsedPrices[price] = name
		}

		if valid {
			for price, name := range priceToName {
				if parsedName, exists := parsedPrices[price]; !exists || parsedName != name {
					valid = false
					break
				}
			}

			for price, cnt := range priceCount {
				if cnt == 1 {
					continue
				}
				if _, exists := parsedPrices[price]; !exists {
					valid = false
					break
				}
			}
		}

		if valid {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isValidNameFast(name string) bool {
	if len(name) == 0 || len(name) > 10 {
		return false
	}
	for i := 0; i < len(name); i++ {
		if c := name[i]; c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func isValidPriceStrFast(priceStr string) bool {
	if len(priceStr) == 0 || (len(priceStr) > 1 && priceStr[0] == '0') {
		return false
	}
	for i := 0; i < len(priceStr); i++ {
		if c := priceStr[i]; c < '0' || c > '9' {
			return false
		}
	}
	return true
}
