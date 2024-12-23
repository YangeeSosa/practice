package leetcode

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	startI := 0
	sign := 1
	if s[0] == '-' {
		sign = -1
		startI++
	} else if s[0] == '+' {
		startI++
	}

	result := 0
	for i := startI; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num := int(s[i] - '0')
		if result > (math.MaxInt32-num)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + num
	}
	return result * sign
}
