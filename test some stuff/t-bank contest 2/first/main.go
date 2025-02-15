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

	indexR := -1
	indexM := -1

	for i, char := range s {
		if char == 'R' {
			indexR = i
		} else if char == 'M' {
			indexM = i
		}
	}

	if indexR < indexM {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
