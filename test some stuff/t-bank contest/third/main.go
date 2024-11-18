package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()

	scanner.Scan()
	required := scanner.Text()

	scanner.Scan()
	var k int
	fmt.Sscanf(scanner.Text(), "%d", &k)

	requiredCount := make(map[rune]int)
	for _, c := range required {
		requiredCount[c]++
	}

	windowCount := make(map[rune]int)
	uniqueCount := 0
	left, bestStart, bestLen := 0, -1, 0

	for right, r := range s {
		windowCount[r]++
		if count, ok := requiredCount[r]; ok && windowCount[r] == count {
			uniqueCount++
		}

		for uniqueCount == len(requiredCount) {
			if right-left+1 <= k && right-left+1 > bestLen {
				bestStart, bestLen = left, right-left+1
			}

			leftChar := rune(s[left])
			windowCount[leftChar]--
			if count, ok := requiredCount[leftChar]; ok && windowCount[leftChar] < count {
				uniqueCount--
			}
			left++
		}
	}

	if bestStart == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(s[bestStart : bestStart+bestLen])
	}
}
