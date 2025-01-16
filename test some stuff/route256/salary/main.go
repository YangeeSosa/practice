package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func maxSalaryAfterRemovingDigit(s string) string {
	n := len(s)
	if n == 1 {
		return "0"
	}

	for i := 0; i < n-1; i++ {
		if s[i] < s[i+1] {
			return s[:i] + s[i+1:]
		}
	}

	return s[:n-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		result := maxSalaryAfterRemovingDigit(s)
		fmt.Fprintln(writer, result)
	}
}
